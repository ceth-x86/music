package db

import "github.com/jinzhu/gorm"

type Track struct {
	gorm.Model
	PlaylistId uint `gorm:"TYPE:integer REFERENCES playlists;unique_index:idx_playlist_track_id"`
	Name       string
	TrackId    string `gotm:"unique_index:idx_playlist_track_id"`
	// Artist           string
	// Album            string
}
