package engine

import (
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

	playlist.Name = servicePlaylist.Name
	playlist.Description = servicePlaylist.Description
	_, err = e.DataRepository.PlaylistRepository.Update(playlistId, playlist)
	if err != nil {
		logger.With(zap.Error(err)).Error("не удалось обновить плейлист")
	}

	for _, track := range tracks {
		track.PlaylistId = playlistId
		_, err = e.DataRepository.TrackRepository.Store(track)
		if err != nil {
			logger.With(zap.Error(err)).Errorw("не удалось сохранить трек")
		}
	}
}
