package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Playlist struct {
	gorm.Model
	Service     string `gorm:"index:idx_service_playlist_id"`
	PlaylistId  string `gorm:"index:idx_service_playlist_id"`
	TrackCount  *uint
	LastChanged *time.Time
}
