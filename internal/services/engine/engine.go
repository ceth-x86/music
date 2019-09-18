package engine

import (
	"time"

	"github.com/demas/music/internal/models/core"

	"github.com/demas/music/internal/services/datastore/repository"
	"github.com/demas/music/internal/services/musicservices"
	"go.uber.org/zap"
)

type Engine struct {
	DataRepository *repository.Repository
}

func NewEngine(dataRepository *repository.Repository) *Engine {
	return &Engine{DataRepository: dataRepository}
}

func (e *Engine) DownloadPlaylist(playlistId uint) {

	logger := zap.NewExample().Sugar()
	defer func() {
		_ = logger.Sync()
	}()

	playlist, err := e.DataRepository.PlaylistRepository.GetById(playlistId)
	if err != nil {
		logger.Errorw("Cannot find playlist",
			"PlaylistId", playlistId)
		return
	}

	servicePlaylist, tracks, err :=
		musicservices.NewMusicService(playlist.Service).DownloadPlaylist(playlist.PlaylistId)
	if err != nil {
		logger.With(zap.Error(err)).Error(err)
	}

	var playlistWasUpdated = false
	for _, track := range tracks {
		track.PlaylistId = playlistId

		_, err := e.DataRepository.TrackRepository.GetByPlaylistAndTrackId(playlistId, track.TrackId)
		if err != nil {

			// TODO: refactore
			artist, err := e.DataRepository.ArtistRepository.GetByArtistId(track.ServiceArtistId)
			if err != nil {
				artist, err = e.DataRepository.ArtistRepository.Store(&core.Artist{ArtistId: track.ServiceArtistId})
				if err != nil {
					logger.With(zap.Error(err)).Errorw("не удалось сохранить исполнителя",
						"Track.ServiceId", track.TrackId,
						"Track.PlaylistId", track.PlaylistId,
						"Track.ArtistId", track.ServiceArtistId)
					// TODO: refactore
					panic(err)
				}
			}

			album, err := e.DataRepository.AlbumRepository.GetByAlbumId(track.ServiceAlbumId)
			if err != nil {
				album, err = e.DataRepository.AlbumRepository.Store(&core.Album{AlbumId: track.ServiceAlbumId})
				if err != nil {
					logger.With(zap.Error(err)).Errorw("не удалось сохранить альбом",
						"Track.ServiceId", track.TrackId,
						"Track.PlaylistId", track.PlaylistId,
						"Track.AlbumId", track.ServiceAlbumId)
					// TODO: refactore
					panic(err)
				}
			}

			track.ArtistId = artist.Id
			track.AlbumId = album.Id

			_, err = e.DataRepository.TrackRepository.Store(track)
			if err != nil {
				logger.With(zap.Error(err)).Errorw("не удалось сохранить трек")
			}
			playlistWasUpdated = true
		}

	}

	playlist.Name = servicePlaylist.Name
	playlist.Description = servicePlaylist.Description
	if playlistWasUpdated {
		t := time.Now()
		playlist.LastChanged = &t
	}

	_, err = e.DataRepository.PlaylistRepository.Update(playlistId, playlist)
	if err != nil {
		logger.With(zap.Error(err)).Error("не удалось обновить плейлист")
	}
}
