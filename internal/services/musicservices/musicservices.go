package musicservices

import (
	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/models/enums"
	spotify2 "github.com/demas/music/internal/services/musicservices/spotify"
	yandex2 "github.com/demas/music/internal/services/musicservices/yandex"
	"go.uber.org/zap"
)

type IMusicService interface {
	DownloadPlaylist(playlistId string) (*core.Playlist, []*core.Track, error)
}

type IMasterDataRepository interface {
	DownloadAlbum(albumId string) (*core.Album, error)
	DownloadArtist(artistId string) (*core.Artist, error)
	SearchAlbum(artist string, album string) []*core.Album
}

func NewMusicRepository() IMasterDataRepository {
	return spotify2.NewSpotify()
}

func NewMusicService(service uint) IMusicService {

	logger := zap.NewExample().Sugar()
	defer func() {
		_ = logger.Sync()
	}()

	switch service {
	case uint(enums.MusicServiceSpotify):
		return spotify2.NewSpotify()
	case uint(enums.MusicServiceYandex):
		return yandex2.NewYandex()
	default:
		logger.Errorw("not implemented music service",
			"Service", service)
		return nil
	}
}
