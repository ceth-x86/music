package datastore

import (
	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/models/db"
	"github.com/jinzhu/gorm"
)

type IReleaseRepository interface {
	Store(p *core.Release) (*core.Release, error)
}

type ReleaseRepository struct {
	db *gorm.DB
}

func NewReleaseRepository(dbHandler *gorm.DB) *ReleaseRepository {
	return &ReleaseRepository{
		db: dbHandler,
	}
}

func setReleaseFields(r db.Release, release *core.Release) db.Release {
	r.AlbumId = release.AlbumId
	r.SyncDate = release.SyncDate
	return r
}

func (r *ReleaseRepository) Store(rl *core.Release) (*core.Release, error) {

	release := setReleaseFields(db.Release{}, rl)
	if err := r.db.Create(&release).Error; err != nil {
		return nil, err
	}
	rl.Id = release.ID
	return rl, nil
}
