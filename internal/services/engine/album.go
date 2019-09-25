package engine

import (
	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/services/musicservices"
)

// bool: is it new album, or existing one
func (e *Engine) returnOrCreateAlbum(musicService musicservices.IMusicService, serviceAlbumId string, artistId uint) (*core.Album, bool, error) {

	album, err := e.DataRepository.AlbumRepository.GetByAlbumId(serviceAlbumId)
	if err != nil {

		album, err = musicService.DownloadAlbum(serviceAlbumId)
		if err != nil {
			return nil, false, &DownloadError{Cause: err}
		}

		album.ArtistId = artistId
		album, err = e.DataRepository.AlbumRepository.Store(album)
		if err != nil {
			return nil, false, &StoreError{Cause: err}
		}

		return album, true, nil
	}

	return album, false, nil
}
