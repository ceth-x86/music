package musicservices

import (
	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/models/enums"
	spotify2 "github.com/demas/music/internal/services/musicservices/spotify"
	"go.uber.org/zap"
)

type IMusicService interface {
	DownloadPlaylist(playlistId string) (*core.Playlist, []*core.Track, error)
}

func NewMusicService(service uint) IMusicService {

	logger := zap.NewExample().Sugar()
	defer func() {
		_ = logger.Sync()
	}()

	switch service {
	case uint(enums.MusicServiceSpotify):
		spotify := spotify2.NewSpotify()
		return spotify
	default:
		logger.Errorw("not implemented music service",
			"Service", service)
		return nil
	}
}
