package datastore

import (
	"errors"

	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/models/db"
	"github.com/jinzhu/gorm"
)

type IAlbumRepository interface {
	GetById(id uint) (*core.Album, error)
	GetByAlbumId(albumId string) (*core.Album, error)
	Store(p *core.Album) (*core.Album, error)
}

type AlbumRepository struct {
	db *gorm.DB
}

func NewAlbumRepository(dbHandler *gorm.DB) *AlbumRepository {
	return &AlbumRepository{
		db: dbHandler,
	}
}

func setAlbumFields(a db.Album, album *core.Album) db.Album {
	a.AlbumId = album.AlbumId
	a.Name = album.Name
	a.Popularity = album.Popularity
	a.ReleaseDate = album.ReleaseDate
	a.AlbumType = album.AlbumType
	return a
}

func (r *AlbumRepository) GetById(id uint) (*core.Album, error) {
	result := &core.Album{}
	r.db.First(&result, id)
	if result.Id == 0 {
		return nil, errors.New("record not found")
	}
	return result, nil
}

func (r *AlbumRepository) GetByAlbumId(albumId string) (*core.Album, error) {
	result := &core.Album{}
	r.db.Where("album_id = ?", albumId).First(&result)
	if result.Id == 0 {
		return nil, errors.New("record not found")
	}
	return result, nil
}

func (r *AlbumRepository) Store(a *core.Album) (*core.Album, error) {

	album := setAlbumFields(db.Album{}, a)
	if err := r.db.Create(&album).Error; err != nil {
		return nil, err
	}
	a.Id = album.ID
	return a, nil
}
