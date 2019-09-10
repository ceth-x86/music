package datastore

import (
	"errors"

	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/models/db"
	"github.com/jinzhu/gorm"
)

type IPlaylistRepository interface {
	Fetch() []*core.Playlist
	GetById(id uint) (*core.Playlist, error)
	Store(p *core.Playlist) (*core.Playlist, error)
	Update(id uint, p *core.Playlist) (*core.Playlist, error)
}

type PlaylistRepository struct {
	db *gorm.DB
}

func NewPlaylistRepository(dbHandler *gorm.DB) *PlaylistRepository {
	return &PlaylistRepository{
		db: dbHandler,
	}
}

func setPlaylistFields(p db.Playlist, playlist *core.Playlist) db.Playlist {
	p.Service = playlist.Service
	p.PlaylistId = playlist.PlaylistId
	p.Name = playlist.Name
	p.Description = playlist.Description
	p.TrackCount = playlist.TrackCount
	p.LastChanged = playlist.LastChanged
	return p
}

func (r *PlaylistRepository) Fetch() []*core.Playlist {

	var result []*core.Playlist
	r.db.Where("deleted_at is null").Find(&result)
	return result
}

func (r *PlaylistRepository) GetById(id uint) (*core.Playlist, error) {
	result := &core.Playlist{}
	r.db.First(&result, id)
	if result.Id == 0 {
		return nil, errors.New("record not found")
	}
	return result, nil
}

func (r *PlaylistRepository) Store(p *core.Playlist) (*core.Playlist, error) {

	playlist := setPlaylistFields(db.Playlist{}, p)
	if err := r.db.Create(&playlist).Error; err != nil {
		return nil, err
	}
	p.Id = playlist.ID
	return p, nil
}

func (r *PlaylistRepository) Update(id uint, p *core.Playlist) (*core.Playlist, error) {

	playlist := db.Playlist{}
	r.db.First(&playlist, id)
	playlist = setPlaylistFields(playlist, p)
	if err := r.db.Save(&playlist).Error; err != nil {
		return nil, err
	}
	return p, nil
}
