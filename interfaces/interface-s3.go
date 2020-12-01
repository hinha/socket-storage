package interfaces

import (
	"bytes"
	"github.com/atreugo/websocket"
	"github.com/fasthttp-contrib/render"
	"github.com/savsgio/atreugo/v11"
	"socket-storage/models"
	file_stream "socket-storage/py-rpc/proto"
)

type StorageS3Controller interface {
	UploadFile(ws *websocket.Conn) error
	LogFile(ws *websocket.Conn) error
	LogFileRender(ctx *atreugo.RequestCtx) error
	UploadFileRender(ctx *atreugo.RequestCtx) error
}

type StorageS3Persistence interface {
	DuplicatesFile(Filename string, UserID int) error
	InsertFileData(dataModel *models.DataModel) (int, error)
}

type StorageS3Service interface {
	UploadFile(ws *websocket.Conn) error
	LogFile(ws *websocket.Conn) error
	LogFileRender(ctx *atreugo.RequestCtx, tpl *render.Render) error
	UploadFileRender(ctx *atreugo.RequestCtx, tpl *render.Render) error
}

type StorageS3Repository interface {
	InsertFileData(modelData *models.DataModel) (int, error)
	FilterDuplicatesFile(FileEncrypt string, UserID int) error
	HashFileMD5(fileReader *bytes.Reader) (string, error)
	PutObject(fileReader *bytes.Reader, message []byte, fileName, FileEncrypt string) error
	PutFileStreamProto(request *file_stream.InputFrame) (*file_stream.OutputFrame, error)
	ResponseFileStreamProto(message []byte, fileName, fileExt string, userID int) (string, *file_stream.OutputFrame, error)
}
