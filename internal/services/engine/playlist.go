package engine

import (
	"time"

	"github.com/demas/music/internal/services/musicservices"

	"github.com/demas/music/internal/models/core"
	"go.uber.org/zap"
)

func (e *Engine) processDownloadedTrack(track *core.Track, musicService musicservices.IMusicService) (wasSaved bool) {

	logger := zap.NewExample().Sugar()
	defer func() {
		_ = logger.Sync()
	}()

	_, err := e.DataRepository.TrackRepository.GetByPlaylistAndTrackId(track.PlaylistId, track.TrackId)
	if err != nil {

		// artist
		artist, err := e.returnOrCreateArtist(musicService, track.ServiceArtistId)
		if err != nil {
			switch e := err.(type) {
			case *DownloadError:
				logger.With(zap.Error(e)).Errorw("не удалось получить исполнителя на музыкальном сервисе",
					"Track.ServiceId", track.TrackId,
					"Track.PlaylistId", track.PlaylistId,
					"Track.ArtistId", track.ServiceArtistId)
			case *StoreError:
				logger.With(zap.Error(err)).Errorw("не удалось сохранить исполнителя",
					"Track.ServiceId", track.TrackId,
					"Track.PlaylistId", track.PlaylistId,
					"Track.ArtistId", track.ServiceArtistId)
			}
			return false
		}

		// album
		album, newAlbum, err := e.returnOrCreateAlbum(musicService, track.ServiceAlbumId, artist.Id)
		if err != nil {
			switch e := err.(type) {
			case *DownloadError:
				logger.With(zap.Error(e)).Errorw("не удалось получить альбом на музыкальном сервисе",
					"Track.ServiceId", track.TrackId,
					"Track.PlaylistId", track.PlaylistId,
					"Track.AlbumId", track.ServiceAlbumId)
			case *StoreError:
				logger.With(zap.Error(err)).Errorw("не удалось сохранить альбом",
					"Track.ServiceId", track.TrackId,
					"Track.PlaylistId", track.PlaylistId,
					"Track.AlbumId", track.ServiceAlbumId)
			}
			return false
		}
		album.ArtistId = artist.Id

		if newAlbum && isItNewRelease(album.ReleaseDate) {
			_, err := e.DataRepository.ReleaseRepository.Store(&core.Release{
				AlbumId:  album.Id,
				SyncDate: time.Now(),
			})

			if err != nil {
				logger.With(zap.Error(err)).Errorw("не удалось сохранить релиз",
					"Track.ServiceId", track.TrackId,
					"Track.PlaylistId", track.PlaylistId,
					"Track.AlbumId", track.ServiceAlbumId)
				return
			}
		}

		track.ArtistId = artist.Id
		track.AlbumId = album.Id

		_, err = e.DataRepository.TrackRepository.Store(track)
		if err != nil {
			logger.With(zap.Error(err)).Errorw("не удалось сохранить трек")
		}
		return true
	}
	return false
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
		playlistWasUpdated = e.processDownloadedTrack(track, musicService)
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
