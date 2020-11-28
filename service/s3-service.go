package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/atreugo/websocket"
	"github.com/savsgio/atreugo/v11"
	"socket-storage/interfaces"
	"socket-storage/vo"
	"strings"
)

type StorageS3Service struct {
	storageRepository interfaces.StorageS3Repository
}

const MaxSizeFile = 1024 * 1024 * 85 // 25mb in files

func (c *StorageS3Service) Test(ws *websocket.Conn) error {
	var allowedExt bool
	//detector := detector.New() delete

	for {
		_, message, err := ws.ReadMessage()
		body := bytes.NewReader(message)
		if err != nil {
			return err
		}

		uploadFile := new(vo.UploadFile)
		err = json.Unmarshal(message, uploadFile)
		if err != nil {
			ws.WriteJSON(atreugo.JSON{
				"status": "failed",
				"reason": "error payload",
			})
		}
		uploadFile.Type = strings.ToLower(uploadFile.Type) // to lower_case

		fmt.Println(body.Size(), MaxSizeFile)
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

			//if true {} adding logic db Duplicates files

			resp, err := c.storageRepository.PutObject(body, uploadFile.Result, uploadFile.Name, fileExt, "1")
			if err != nil {
				ws.WriteJSON(atreugo.JSON{
					"status": "failed",
					"reason": "something went wrong. can't upload file",
				})
			}
			fmt.Println(resp)
		} else {
			ws.WriteJSON(atreugo.JSON{
				"status": "failed",
				"reason": "file not allowed",
			})
		}
	}
}

func NewStorageS3Service(repository interfaces.StorageS3Repository) interfaces.StorageS3Service {
	return &StorageS3Service{storageRepository: repository}
}
