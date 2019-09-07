package repository

import (
	"github.com/demas/music/internal/services/datastore"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	PlaylistRepository datastore.IPlaylistRepository
	TrackRepository    datastore.ITrackRepository
}

func NewRepository(dbHandler *gorm.DB) *Repository {
	return &Repository{
		PlaylistRepository: datastore.NewPlaylistRepository(dbHandler),
		TrackRepository:    datastore.NewTrackRepository(dbHandler),
	}
}
