package musicservices

import (
	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/models/enums"
	spotify2 "github.com/demas/music/internal/services/musicservices/spotify"
	yandex2 "github.com/demas/music/internal/services/musicservices/yandex"
	"go.uber.org/zap"
)

// any playlists provider
type IMusicService interface {
	DownloadPlaylist(playlistId string) (*core.Playlist, []*core.Track, error)
}

// master data repository - we need to have one source of truth
type IMasterDataRepository interface {
	SearchAlbum(artist string, album string) ([]*core.Album, error) // search album by artist and album name (we get them from playlists)
	DownloadArtist(artistId string) (*core.Artist, error)           // get artist by music service artist ID
	DownloadAlbum(albumId string) (*core.Album, error)              // get album by music service album ID
}

// I have only one MasterData provider (may be in the future I'll add LastFM)
func NewMusicRepository() IMasterDataRepository {
	return spotify2.NewSpotify()
}

// I have two playlists providers: Spotify and Yandex.Music
// I am planning to add Deezer
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
