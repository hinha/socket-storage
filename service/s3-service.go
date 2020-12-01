package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/atreugo/websocket"
	"github.com/fasthttp-contrib/render"
	"github.com/getsentry/sentry-go"
	"github.com/hpcloud/tail"
	"github.com/savsgio/atreugo/v11"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"socket-storage/interfaces"
	"socket-storage/models"
	"socket-storage/vo"
	"strings"
	"time"
)

type StorageS3Service struct {
	storageRepository interfaces.StorageS3Repository
}

var (
	filename string
)

const (
	// Time allowed to write the file to the client.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Poll file for changes with this period.
	filePeriod = 10 * time.Second

	// TextMessage denotes a text data message. The text message payload is
	// interpreted as UTF-8 encoded text data.
	TextMessage = 1

	// BinaryMessage denotes a binary data message.
	BinaryMessage = 2

	// CloseMessage denotes a close control message. The optional message
	// payload contains a numeric code and text. Use the FormatCloseMessage
	// function to format a close message payload.
	CloseMessage = 8

	// PingMessage denotes a ping control message. The optional message payload
	// is UTF-8 encoded text.
	PingMessage = 9

	// PongMessage denotes a pong control message. The optional message payload
	// is UTF-8 encoded text.
	PongMessage = 10
)

const MaxSizeFile = 1024 * 1024 * 85 // 25mb in files

func (c *StorageS3Service) UploadFile(ws *websocket.Conn) error {
	var allowedExt bool
	//detector := detector.New() delete

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			sentry.CaptureException(err)
			return err
		}
		body := bytes.NewReader(message)

		uploadFile := new(vo.UploadFile)
		err = json.Unmarshal(message, uploadFile)
		if err != nil {
			ws.WriteJSON(atreugo.JSON{
				"status": "failed",
				"reason": "error payload",
			})
		}
		uploadFile.Type = strings.ToLower(uploadFile.Type) // to lower_case

		if body.Size() > MaxSizeFile {
			ws.WriteJSON(atreugo.JSON{
				"status": "failed",
				"reason": "file size must less than 15mb",
			})
			return nil
		}

		switch uploadFile.Type {
		case "text/csv":
			allowedExt = true
		case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
			allowedExt = true
		case "application/vnd.ms-excel":
			allowedExt = true
		default:
			allowedExt = false
		}

		if allowedExt {
			// split file extension
			fSep := strings.Split(uploadFile.Name, ".")
			fileExt := strings.ToLower(fSep[len(fSep)-1:][0])

			ByteCount, FileStream, err := c.storageRepository.ResponseFileStreamProto(message, uploadFile.Name, fileExt, uploadFile.UserID)
			if err != nil {
				ws.WriteJSON(atreugo.JSON{
					"status": "failed",
					"reason": "something went wrong. can't upload file",
					"error":  fmt.Sprintf("%v", err.Error()),
				})

			} else {

				//if true {} adding logic db Duplicates files
				err = c.storageRepository.FilterDuplicatesFile(FileStream.FileEncrypt, uploadFile.UserID)
				if err == nil { // because is duplicate
					ws.WriteJSON(atreugo.JSON{
						"status": "failed",
						"reason": "file has been added or not found the user id",
					})

					return nil
				}

				err = c.storageRepository.PutObject(body, uploadFile.Result, uploadFile.Name, FileStream.FileEncrypt)
				if err != nil {
					ws.WriteJSON(atreugo.JSON{
						"status": "failed",
						"reason": "something went wrong. can't upload file",
					})
					return nil
				}

				dataModel := models.DataModel{
					FileName: fSep[0],
					Format:   fileExt,
					Columns:  FileStream.Cols,
					Rows:     FileStream.Rows,
					FileSize: ByteCount,
					Url:      fmt.Sprintf("https://s3.ap-southeast-1.amazonaws.com/%s/temp/%s", os.Getenv("BUCKET_NAME"), FileStream.FileEncrypt),
					FileEnc:  FileStream.FileEncrypt,
					UserID:   uploadFile.UserID,
					Status:   "",
				}
				dataID, err := c.storageRepository.InsertFileData(&dataModel)
				if err != nil {
					ws.WriteJSON(atreugo.JSON{
						"status": "failed",
						"reason": "something went wrong. can't upload file model database",
					})
					return nil
				}

				splitHash := strings.Split(FileStream.FileEncrypt, ".")
				ws.WriteJSON(atreugo.JSON{
					"status":   "success",
					"filename": splitHash[0],
					"user_id":  uploadFile.UserID,
					"id":       dataID,
				})
			}

		} else {
			ws.WriteJSON(atreugo.JSON{
				"status": "failed",
				"reason": "file not allowed",
			})
		}
	}
}

func (c *StorageS3Service) LogFile(ws *websocket.Conn) error {
	var lastMod time.Time

	go writer(ws, lastMod)
	ws.PingHandler()
	reader(ws)
	return nil
}

func (c *StorageS3Service) LogFileRender(ctx *atreugo.RequestCtx, render *render.Render) error {
	path, _ := os.Getwd()

	y, m, _ := time.Now().Date()

	filename = fmt.Sprintf("%s/temp/log/%d%d_%s", path, int(m), y, "info.log")
	_, dataLog, lastMod, err := readFileIfModified(time.Time{})
	if err != nil {
		lastMod = time.Unix(0, 0)
	}
	//fmt.Println(string(p))

	data := atreugo.JSON{
		"title":    "index",
		"host":     string(ctx.Host()),
		"data":     dataLog,
		"last_mod": lastMod,
	}
	return render.HTML(ctx.RequestCtx, http.StatusOK, "logging", data)
}

func (c *StorageS3Service) UploadFileRender(ctx *atreugo.RequestCtx, render *render.Render) error {
	return render.HTML(ctx.RequestCtx, http.StatusOK, "index", "")
}

func NewStorageS3Service(repository interfaces.StorageS3Repository) interfaces.StorageS3Service {
	return &StorageS3Service{storageRepository: repository}
}

func readFileIfModified(lastMod time.Time) ([]byte, []vo.Logger, time.Time, error) {
	var logStruct vo.Logger
	var dataLog []vo.Logger

	fi, err := os.Stat(filename)
	if err != nil {
		return nil, nil, lastMod, err
	}
	if !fi.ModTime().After(lastMod) {
		return nil, nil, lastMod, nil
	}

	t, err := tail.TailFile(filename, tail.Config{Follow: false})
	if err != nil {
		return nil, nil, lastMod, err
	}
	for line := range t.Lines {
		json.Unmarshal([]byte(line.Text), &logStruct)
		dataLog = append(dataLog, logStruct)
	}

	p, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, fi.ModTime(), err
	}
	return p, dataLog, fi.ModTime(), nil
}

func reader(ws *websocket.Conn) {
	defer func() {
		ws.Close()
	}()
	fmt.Println("reader")
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
	}
}

func writer(ws *websocket.Conn, lastMod time.Time) {
	var dataLog []vo.Logger
	lastError := ""
	pingTicker := time.NewTicker(pingPeriod)
	fileTicker := time.NewTicker(filePeriod)
	defer func() {
		fmt.Println("opss")
		pingTicker.Stop()
		fileTicker.Stop()
		ws.Close()
	}()

	if err := ws.Conn.PingHandler(); err == nil {
		fmt.Println("got channel close")
	}

	for {
		fmt.Println("rinn")

		select {
		case <-fileTicker.C:
			var p []byte
			var err error

			p, dataLog, lastMod, err = readFileIfModified(lastMod)
			if err != nil {
				if s := err.Error(); s != lastError {
					lastError = s
					p = []byte(lastError)
				}
			} else {
				lastError = ""
			}

			reqBodyBytes := new(bytes.Buffer)
			json.NewEncoder(reqBodyBytes).Encode(dataLog)

			if p != nil {
				ws.SetWriteDeadline(time.Now().Add(writeWait))
				if err := ws.WriteMessage(TextMessage, reqBodyBytes.Bytes()); err != nil {
					return
				}
			}
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(PingMessage, []byte{}); err != nil {
				return
			}
		}

	}
}
