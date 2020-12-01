package persistence

import (
	"socket-storage/interfaces"
	"socket-storage/models"
	"socket-storage/platform"
)

type StorageS3Persistence struct {
	db *platform.DatabaseORM
}

func (s *StorageS3Persistence) DuplicatesFile(FileEncrypt string, UserID int) error {
	result := new(models.DataModel)
	err := s.db.Master.Table("data").Where("file_enc = ? AND user_id = ?", FileEncrypt, UserID).First(result).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *StorageS3Persistence) InsertFileData(dataModel *models.DataModel) (int, error) {
	result := s.db.Master.Table("data").Create(dataModel)
	if result.Error != nil {
		return 0, result.Error
	}
	return dataModel.UserID, nil
}

func NewStorageS3Service(db *platform.DatabaseORM) interfaces.StorageS3Persistence {
	return &StorageS3Persistence{db: db}
}
