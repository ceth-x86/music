package engine

import (
	"fmt"
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

	musicService := musicservices.NewMusicService(playlist.Service)

	servicePlaylist, tracks, err := musicService.DownloadPlaylist(playlist.PlaylistId)
	if err != nil {
		logger.With(zap.Error(err)).Error(err)
	}

	var playlistWasUpdated = false
	for _, track := range tracks {
		track.PlaylistId = playlistId

		_, err := e.DataRepository.TrackRepository.GetByPlaylistAndTrackId(playlistId, track.TrackId)
		if err != nil {

			// artist
			artist, err := e.DataRepository.ArtistRepository.GetByArtistId(track.ServiceArtistId)
			if err != nil {

				artist, err = musicService.DownloadArtist(track.ServiceArtistId)
				if err != nil {
					logger.With(zap.Error(err)).Errorw("не удалось получить исполнителя на музыкальном сервисе",
						"Track.ServiceId", track.TrackId,
						"Track.PlaylistId", track.PlaylistId,
						"Track.ArtistId", track.ServiceArtistId)
					continue
				}

				artist, err = e.DataRepository.ArtistRepository.Store(artist)
				if err != nil {
					logger.With(zap.Error(err)).Errorw("не удалось сохранить исполнителя",
						"Track.ServiceId", track.TrackId,
						"Track.PlaylistId", track.PlaylistId,
						"Track.ArtistId", track.ServiceArtistId)
					continue
				}
			}

			// album
			album, err := e.DataRepository.AlbumRepository.GetByAlbumId(track.ServiceAlbumId)
			if err != nil {

				album, err = musicService.DownloadAlbum(track.ServiceAlbumId)
				if err != nil {
					logger.With(zap.Error(err)).Errorw("не удалось получить альбом на музыкальном сервисе",
						"Track.ServiceId", track.TrackId,
						"Track.PlaylistId", track.PlaylistId,
						"Track.AlbumId", track.ServiceAlbumId)
					continue
				}

				album, err = e.DataRepository.AlbumRepository.Store(album)
				if err != nil {
					logger.With(zap.Error(err)).Errorw("не удалось сохранить альбом",
						"Track.ServiceId", track.TrackId,
						"Track.PlaylistId", track.PlaylistId,
						"Track.AlbumId", track.ServiceAlbumId)
					continue
				}

				// если вышло не позже месяца, то считаем, что это новинка
				d := time.Now().Sub(album.ReleaseDate).Hours()
				fmt.Println(d)
				if time.Now().Sub(album.ReleaseDate).Hours() < 24*30 {
					_, err := e.DataRepository.ReleaseRepository.Store(&core.Release{
						AlbumId:  album.Id,
						SyncDate: time.Now(),
					})

					if err != nil {
						logger.With(zap.Error(err)).Errorw("не удалось сохранить релиз",
							"Track.ServiceId", track.TrackId,
							"Track.PlaylistId", track.PlaylistId,
							"Track.AlbumId", track.ServiceAlbumId)
						continue
					}
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
