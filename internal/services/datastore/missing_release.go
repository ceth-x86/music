package datastore

import (
	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/models/db"
	"github.com/jinzhu/gorm"
)

type IMissingReleaseRepository interface {
	Fetch() []*core.MissingRelease
	FindByAlbumAndArtist(albumName string, artistName string) []*core.MissingRelease
	Store(r *core.MissingRelease) (*core.MissingRelease, error)
}

type MissingReleaseRepository struct {
	db *gorm.DB
}

func NewMissingReleaseRepository(dbHandler *gorm.DB) *MissingReleaseRepository {
	return &MissingReleaseRepository{
		db: dbHandler,
	}
}

func setMissingReleaseFields(r db.MissingRelease, release *core.MissingRelease) db.MissingRelease {
	r.AlbumName = release.AlbumName
	r.ArtistName = release.ArtistName
	return r
}

func (r *MissingReleaseRepository) Fetch() []*core.MissingRelease {

	var result []*core.MissingRelease
	r.db.Where("deleted_at is null").Find(&result)
	return result
}

func (r *MissingReleaseRepository) FindByAlbumAndArtist(albumName string, artistName string) []*core.MissingRelease {

	var result []*core.MissingRelease
	r.db.Where("album_name = ? and artist_name = ? and deleted_at is null", albumName, artistName).Find(&result)
	return result
}

func (r *MissingReleaseRepository) Store(mr *core.MissingRelease) (*core.MissingRelease, error) {

	missingRelease := setMissingReleaseFields(db.MissingRelease{}, mr)
	if err := r.db.Create(&missingRelease).Error; err != nil {
		return nil, err
	}
	mr.Id = missingRelease.ID
	return mr, nil
}
