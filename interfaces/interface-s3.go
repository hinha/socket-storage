package interfaces

import (
	"bytes"
	"github.com/atreugo/websocket"
)

type StorageS3Controller interface {
	Test(ws *websocket.Conn) error
}

type StorageS3Service interface {
	Test(ws *websocket.Conn) error
}

type StorageS3Repository interface {
	HashFileMD5(fileReader *bytes.Reader) (string, error)
	PutObject(fileReader *bytes.Reader, message []byte, fileName, fileExt string) error
}
