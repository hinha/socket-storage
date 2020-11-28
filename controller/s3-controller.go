package controller

import (
	"github.com/atreugo/websocket"
	"socket-storage/interfaces"
)

type StorageS3Controller struct {
	storageService interfaces.StorageS3Service
}

func (c *StorageS3Controller) Test(ws *websocket.Conn) error {
	return c.storageService.Test(ws)
}

func NewStorageS3Controller(service interfaces.StorageS3Service) interfaces.StorageS3Controller {
	return &StorageS3Controller{storageService: service}
}
