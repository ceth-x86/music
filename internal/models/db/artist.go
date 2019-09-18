package db

import "github.com/jinzhu/gorm"

type Artist struct {
	gorm.Model
	ArtistId string `gorm:"unique_index"`
	Name     string
}
