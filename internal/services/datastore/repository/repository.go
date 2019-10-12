package repository

import (
	"github.com/demas/music/internal/services/datastore"
	"github.com/jackc/pgx/v4"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	AlbumRepository          datastore.IAlbumRepository
	ArtistRepository         datastore.IArtistRepository
	MissingReleaseRepository datastore.IMissingReleaseRepository
	PlaylistRepository       datastore.IPlaylistRepository
	ReleaseRepository        datastore.IReleaseRepository
	TrackRepository          datastore.ITrackRepository
}

func NewRepository(dbHandler *gorm.DB, connection *pgx.Conn) *Repository {
	return &Repository{
		AlbumRepository:          datastore.NewAlbumRepository(dbHandler),
		ArtistRepository:         datastore.NewArtistRepository(dbHandler),
		MissingReleaseRepository: datastore.NewMissingReleaseRepository(dbHandler),
		PlaylistRepository:       datastore.NewPlaylistRepository(dbHandler, connection),
		ReleaseRepository:        datastore.NewReleaseRepository(dbHandler),
		TrackRepository:          datastore.NewTrackRepository(dbHandler),
	}
}
