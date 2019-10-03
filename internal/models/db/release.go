package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Release struct {
	gorm.Model
	AlbumId    uint      `gorm:"TYPE:integer REFERENCES albums"`
	PlaylistId uint      `gorm:"TYPE:integer REFERENCES playlists"`
	SyncDate   time.Time `gorm:"index"`
}
