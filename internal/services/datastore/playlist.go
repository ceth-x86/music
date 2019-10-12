package datastore

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"

	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/models/db"
	"github.com/jinzhu/gorm"
)

type IPlaylistRepository interface {
	Fetch() ([]*core.Playlist, error)
	GetById(id uint) (*core.Playlist, error)
	Store(p *core.Playlist) (*core.Playlist, error)
	Update(id uint, p *core.Playlist) (*core.Playlist, error)
}

type PlaylistRepository struct {
	db   *gorm.DB
	conn *pgx.Conn
}

func NewPlaylistRepository(dbHandler *gorm.DB, conn *pgx.Conn) *PlaylistRepository {
	return &PlaylistRepository{
		db:   dbHandler,
		conn: conn,
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

func (r *PlaylistRepository) Fetch() ([]*core.Playlist, error) {

	var sql = "select id, service, playlist_id, name, description, last_changed from playlists where deleted_at is null order by id"
	var result []*core.Playlist

	rows, err := r.conn.Query(context.Background(), sql)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {

		var (
			id          uint
			service     uint
			playlistId  string
			name        string
			description string
			lastChanged *time.Time
		)

		err = rows.Scan(&id, &service, &playlistId, &name, &description, &lastChanged)
		if err != nil {
			return result, err
		}

		playlist := new(core.Playlist)
		playlist.Id = id
		playlist.Service = service
		playlist.PlaylistId = playlistId
		playlist.Name = name
		playlist.Description = description
		playlist.LastChanged = lastChanged
		result = append(result, playlist)
	}

	return result, nil
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
