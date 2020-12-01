package models

import "time"

type DataModel struct {
	ID        int       `gorm:"primary_key;auto_increment;not_null" json:"id"`
	FileName  string    `gorm:"size:64;column:filename" json:"file_name"`
	Format    string    `gorm:"size:64" json:"format"`
	Columns   int64     `gorm:"type:int;column:columns" json:"columns"`
	Rows      int64     `gorm:"type:int;column:rows" json:"rows"`
	FileSize  string    `gorm:"size:64" json:"file_size"`
	Url       string    `gorm:"size:254" json:"url"`
	FileEnc   string    `gorm:"size:64" json:"file_enc"`
	UserID    int       `gorm:"type:int;column:user_id" json:"user_id"`
	Status    string    `gorm:"size:20" json:"status"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp"`
}
