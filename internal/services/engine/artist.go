package engine

import (
	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/services/musicservices"
)

func (e *Engine) returnOrCreateArtist(musicService musicservices.IMusicRepository, serviceArtistId string) (*core.Artist, error) {

	artist, err := e.DataRepository.ArtistRepository.GetByArtistId(serviceArtistId)
	if err != nil {
		artist, err = musicService.DownloadArtist(serviceArtistId)
		if err != nil {
			return nil, &DownloadError{Cause: err}
		}

		artist, err = e.DataRepository.ArtistRepository.Store(artist)
		if err != nil {
			return nil, &StoreError{Cause: err}
		}
	}

	return artist, nil
}
