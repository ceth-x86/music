package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Album struct {
	gorm.Model
	AlbumId     string `gorm:"unique_index"`
	Name        string
	ArtistId    uint `gorm:"TYPE:integer REFERENCES artists"`
	Popularity  int
	AlbumType   string
	ReleaseDate time.Time
}
