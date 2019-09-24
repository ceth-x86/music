package engine

import (
	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/services/musicservices"
)

func (e *Engine) returnOrCreateArtist(musicService musicservices.IMusicService, serviceArtistId string) (*core.Artist, error) {

	artist, err := e.DataRepository.ArtistRepository.GetByArtistId(serviceArtistId)
	if err != nil {
		artist, err = musicService.DownloadArtist(serviceArtistId)
		if err != nil {
			return nil, NewDownloadError(err)
		}

		artist, err = e.DataRepository.ArtistRepository.Store(artist)
		if err != nil {
			return nil, NewStoreError(err)
		}
	}

	return artist, nil
}
