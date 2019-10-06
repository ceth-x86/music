package datastore

import (
	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/models/db"
	"github.com/jinzhu/gorm"
)

type IReleaseRepository interface {
	Fetch() []*core.Release
	GetByAlbumType(albumType string) []*core.Release
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
	r.PlaylistId = release.PlaylistId
	return r
}

func (r *ReleaseRepository) Fetch() []*core.Release {

	var result []*core.Release

	r.db.Select("releases.*, " +
		"albums.name as album_name, albums.album_type as album_type, albums.release_date, " +
		"artists.name as artist_name, artists.genres, " +
		"playlists.name as playlist_name").
		Table("releases").
		Joins("JOIN albums ON releases.album_id = albums.id").
		Joins("JOIN artists ON albums.artist_id = artists.id").
		Joins("JOIN playlists ON releases.playlist_id = playlists.id").
		Where("releases.deleted_at is null").
		Order("albums.release_date, releases.sync_date").
		Find(&result)

	return result
}

func (r *ReleaseRepository) GetByAlbumType(albumType string) []*core.Release {

	var result []*core.Release

	r.db.Select("releases.*, "+
		"albums.name as album_name, albums.album_type as album_type, albums.release_date, "+
		"artists.name as artist_name, artists.genres, "+
		"playlists.name as playlist_name").
		Table("releases").
		Joins("JOIN albums ON releases.album_id = albums.id").
		Joins("JOIN artists ON albums.artist_id = artists.id").
		Joins("JOIN playlists ON releases.playlist_id = playlists.id").
		Where("albums.album_type = ?", albumType).
		Where("releases.deleted_at is null").
		Order("albums.release_date, releases.sync_date").
		Find(&result)

	return result
}

func (r *ReleaseRepository) Store(rl *core.Release) (*core.Release, error) {

	release := setReleaseFields(db.Release{}, rl)
	if err := r.db.Create(&release).Error; err != nil {
		return nil, err
	}
	rl.Id = release.ID
	return rl, nil
}
