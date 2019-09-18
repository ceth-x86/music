package db

import (
	"github.com/jinzhu/gorm"
)

type Album struct {
	gorm.Model
	AlbumId string `gorm:"unique_index"`
	Name    string
}
