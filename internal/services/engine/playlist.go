package engine

import (
	"time"

	"github.com/demas/music/internal/services/musicservices"

	"github.com/demas/music/internal/models/core"
	"go.uber.org/zap"
)

func (e *Engine) processDownloadedTrack(track *core.Track, musicService musicservices.IMusicService) (result *DownloadResult) {

	logger := zap.NewExample().Sugar()
	defer func() {
		_ = logger.Sync()
	}()

	result = &DownloadResult{
		Downloaded: true,
		Album:      0,
		Single:     0,
	}

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
			return DownloadTrackError()
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
			return DownloadTrackError()
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
				return DownloadTrackError()
			}

			if album.AlbumType == "album" {
				result.Album += 1
			} else if album.AlbumType == "single" {
				result.Single += 1
			}
		}

		track.ArtistId = artist.Id
		track.AlbumId = album.Id

		_, err = e.DataRepository.TrackRepository.Store(track)
		if err != nil {
			logger.With(zap.Error(err)).Errorw("не удалось сохранить трек")
		}

		return result
	}

	return DownloadTrackError()
}

func (e *Engine) DownloadPlaylist(playlistId uint) *DownloadResult {

	logger := zap.NewExample().Sugar()
	defer func() {
		_ = logger.Sync()
	}()

	var totalSingles uint = 0
	var totalAlbums uint = 0

	playlist, err := e.DataRepository.PlaylistRepository.GetById(playlistId)
	if err != nil {
		logger.Errorw("Cannot find playlist",
			"PlaylistId", playlistId)
		return DownloadTrackError()
	}

	musicService := musicservices.NewMusicService(playlist.Service)
	servicePlaylist, tracks, err := musicService.DownloadPlaylist(playlist.PlaylistId)
	if err != nil {
		logger.With(zap.Error(err)).Error(err)
	}

	var playlistWasUpdated = false
	for _, track := range tracks {
		track.PlaylistId = playlistId

		trackResult := e.processDownloadedTrack(track, musicService)
		totalAlbums += trackResult.Album
		totalSingles += trackResult.Single
		if trackResult.Downloaded {
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

	return &DownloadResult{Downloaded: true, Album: totalAlbums, Single: totalSingles}
}
