package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Playlist struct {
	gorm.Model
	Service     uint   `gorm:"unique_index:idx_service_playlist_id"`
	PlaylistId  string `gorm:"unique_index:idx_service_playlist_id"`
	Name        string
	Description string
	TrackCount  *uint
	LastChanged *time.Time
}
