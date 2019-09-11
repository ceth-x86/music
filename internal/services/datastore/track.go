package datastore

import (
	"errors"

	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/models/db"
	"github.com/jinzhu/gorm"
)

type ITrackRepository interface {
	GetByPlaylistAndTrackId(playlistId uint, trackId string) (*core.Track, error)
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

// FEATURE: use cache here
func (r *TrackRepository) GetByPlaylistAndTrackId(playlistId uint, trackId string) (*core.Track, error) {
	result := &core.Track{}
	r.db.Where("playlist_id = ? and track_id = ?", playlistId, trackId).First(&result)
	if result.Id == 0 {
		return nil, errors.New("record not found")
	}
	return result, nil
}

func (r *TrackRepository) Store(t *core.Track) (*core.Track, error) {

	track := setTrackFields(db.Track{}, t)
	if err := r.db.Create(&track).Error; err != nil {
		return nil, err
	}
	t.Id = track.ID
	return t, nil
}
