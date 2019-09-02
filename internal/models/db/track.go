package db

import "github.com/jinzhu/gorm"

type Track struct {
	gorm.Model
	PlaylistId uint `gorm:"TYPE:integer REFERENCES playlists"`
	Name       string
	TrackId    string
	// Artist           string
	// Album            string
}
