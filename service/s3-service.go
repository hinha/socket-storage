package service

import (
	"bytes"
	"encoding/json"
	"github.com/atreugo/websocket"
	"github.com/savsgio/atreugo/v11"
	"socket-storage/interfaces"
	"socket-storage/vo"
	"strings"
)

type StorageS3Service struct {
	storageRepository interfaces.StorageS3Repository
}

const MaxSizeFile = 1000000

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

		if body.Size() > MaxSizeFile {
			ws.WriteJSON(atreugo.JSON{
				"status": "failed",
				"reason": "file size must less than 10mb",
			})
		}

		if allowedExt {
			// split file extension
			fSep := strings.Split(uploadFile.Name, ".")
			fileExt := fSep[len(fSep)-1:][0]

			//if true {} adding logic db Duplicates files

			err = c.storageRepository.PutObject(body, uploadFile.Result, uploadFile.Name, fileExt)
			if err != nil {
				ws.WriteJSON(atreugo.JSON{
					"status": "failed",
					"reason": "something went wrong. can't upload file",
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

func NewStorageS3Service(repository interfaces.StorageS3Repository) interfaces.StorageS3Service {
	return &StorageS3Service{storageRepository: repository}
}
