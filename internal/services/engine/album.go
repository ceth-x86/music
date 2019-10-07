package engine

import (
	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/services/musicservices"
)

// bool: is it new album, or existing one
func (e *Engine) returnOrCreateAlbum(musicService musicservices.IMasterDataRepository, serviceAlbumId string, artistId uint) (*core.Album, bool, error) {

	album, err := e.DataRepository.AlbumRepository.GetByAlbumId(serviceAlbumId)
	if err != nil {

		serviceAlbum, err := musicService.DownloadAlbum(serviceAlbumId)
		if err != nil {
			return nil, false, &DownloadError{Cause: err}
		}

		// check album existence with different musicServiceId
		album, err = e.DataRepository.AlbumRepository.GetByArtistIdAndName(artistId, serviceAlbum.Name)
		if err != nil {
			serviceAlbum.ArtistId = artistId
			album, err = e.DataRepository.AlbumRepository.Store(serviceAlbum)
			if err != nil {
				return nil, false, &StoreError{Cause: err}
			}

			return album, true, nil
		}
	}
	return album, false, nil
}
