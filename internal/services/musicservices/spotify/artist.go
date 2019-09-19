package spotify

import (
	"github.com/demas/music/internal/models/core"
	"github.com/zmb3/spotify"
	"go.uber.org/zap"
)

func (s *Spotify) DownloadArtist(artistId string) (*core.Artist, error) {

	logger := zap.NewExample().Sugar()
	spotifyArtist, err := s.Client.GetArtist(spotify.ID(artistId))
	if err != nil {
		logger.With(zap.Error(err)).Error("error getting artist from spotify")
		return nil, err
	}

	artist := &core.Artist{
		Name:     spotifyArtist.Name,
		ArtistId: artistId,
	}

	return artist, nil
}
