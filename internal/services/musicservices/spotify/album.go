package spotify

import (
	"github.com/demas/music/internal/models/core"
	"github.com/zmb3/spotify"
	"go.uber.org/zap"
)

func (s *Spotify) DownloadAlbum(albumId string) (*core.Album, error) {

	logger := zap.NewExample().Sugar()

	spotifyAlbum, err := s.Client.GetAlbum(spotify.ID(albumId))
	if err != nil {
		logger.With(zap.Error(err)).Error("error getting album from spotify")
		return nil, err
	}

	album := &core.Album{
		Name:    spotifyAlbum.Name,
		AlbumId: albumId,
	}

	return album, nil
}
