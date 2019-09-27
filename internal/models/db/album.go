package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Album struct {
	gorm.Model
	AlbumId     string `gorm:"unique_index"`
	Name        string `gorm:"unique_index:idx_artist_id_name"`
	ArtistId    uint   `gorm:"TYPE:integer REFERENCES artists;unique_index:idx_artist_id_name"`
	Popularity  int
	AlbumType   string
	ReleaseDate time.Time
}
