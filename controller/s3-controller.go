package controller

import (
	"github.com/atreugo/websocket"
	"github.com/fasthttp-contrib/render"
	"github.com/savsgio/atreugo/v11"
	"socket-storage/interfaces"
)

type StorageS3Controller struct {
	storageService interfaces.StorageS3Service
	render         *render.Render
}

func (c *StorageS3Controller) UploadFile(ws *websocket.Conn) error {
	return c.storageService.UploadFile(ws)
}

func (c *StorageS3Controller) LogFile(ws *websocket.Conn) error {
	return c.storageService.LogFile(ws)
}

func (c *StorageS3Controller) LogFileRender(ctx *atreugo.RequestCtx) error {
	return c.storageService.LogFileRender(ctx, c.render)
}

func (c *StorageS3Controller) UploadFileRender(ctx *atreugo.RequestCtx) error {
	return c.storageService.UploadFileRender(ctx, c.render)
}

func NewStorageS3Controller(service interfaces.StorageS3Service, render *render.Render) interfaces.StorageS3Controller {
	return &StorageS3Controller{storageService: service, render: render}
}
