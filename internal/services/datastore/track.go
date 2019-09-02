package datastore

import (
	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/models/db"
	"github.com/jinzhu/gorm"
)

type ITrackRepository interface {
	Store(t *core.Track) (*core.Track, error)
}

type TrackRepository struct {
	db *gorm.DB
}

func NewTrackRepository(dbHandler *gorm.DB) *TrackRepository {
	return &TrackRepository{
		db: dbHandler,
	}
}

func setTrackFields(t db.Track, track *core.Track) db.Track {
	t.PlaylistId = track.PlaylistId
	t.Name = track.Name
	t.TrackId = track.TrackId
	return t
}

func (r *TrackRepository) Store(t *core.Track) (*core.Track, error) {

	track := setTrackFields(db.Track{}, t)
	if err := r.db.Create(&track).Error; err != nil {
		return nil, err
	}
	t.Id = track.ID
	return t, nil
}
