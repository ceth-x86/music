package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Release struct {
	gorm.Model
	AlbumId uint `gorm:"TYPE:integer REFERENCES albums"`
	// TODO: index
	SyncDate time.Time
}
