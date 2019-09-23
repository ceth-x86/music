package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Release struct {
	gorm.Model
	AlbumId  uint      `gorm:"TYPE:integer REFERENCES albums"`
	SyncDate time.Time `gorm:"index"`
}
