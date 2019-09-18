package datastore

import (
	"errors"

	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/models/db"
	"github.com/jinzhu/gorm"
)

type IArtistRepository interface {
	GetById(id uint) (*core.Artist, error)
	GetByArtistId(artistId string) (*core.Artist, error)
	Store(p *core.Artist) (*core.Artist, error)
}

type ArtistRepository struct {
	db *gorm.DB
}

func NewArtistRepository(dbHandler *gorm.DB) *ArtistRepository {
	return &ArtistRepository{
		db: dbHandler,
	}
}

func setArtistFields(a db.Artist, artist *core.Artist) db.Artist {
	a.ArtistId = artist.ArtistId
	a.Name = artist.Name
	return a
}

func (r *ArtistRepository) GetById(id uint) (*core.Artist, error) {
	result := &core.Artist{}
	r.db.First(&result, id)
	if result.Id == 0 {
		return nil, errors.New("record not found")
	}
	return result, nil
}

func (r *ArtistRepository) GetByArtistId(artistId string) (*core.Artist, error) {
	result := &core.Artist{}
	r.db.Where("artist_id = ?", artistId).First(&result)
	if result.Id == 0 {
		return nil, errors.New("record not found")
	}
	return result, nil
}

func (r *ArtistRepository) Store(a *core.Artist) (*core.Artist, error) {

	artist := setArtistFields(db.Artist{}, a)
	if err := r.db.Create(&artist).Error; err != nil {
		return nil, err
	}
	a.Id = artist.ID
	return a, nil
}
