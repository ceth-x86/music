package engine

import (
	"errors"
	"fmt"
	"time"

	"github.com/demas/music/internal/services/musicservices"

	"github.com/demas/music/internal/models/core"
	"go.uber.org/zap"
)

// find album and artist IDs using album and artist names
type MasterData struct {
	Engine          *Engine
	Logger          *zap.SugaredLogger
	MusicRepository musicservices.IMasterDataRepository
}

func chooseAlbumFromSearchResults(albums []*core.Album, artistName string, albumName string) *core.Album {

	for _, candidate := range albums {
		if candidate.ReleaseDateString[0:4] == fmt.Sprintf("%v", time.Now().Year()) {
			return candidate
		}
	}
	return nil
}

func (data *MasterData) findArtistAndAlbum(track *core.Track) error {

	missingReleases := data.Engine.DataRepository.MissingReleaseRepository.FindByAlbumAndArtist(track.ServiceAlbumName, track.ServiceArtistName)
	if len(missingReleases) != 0 {
		return errors.New("we already have this track in the list of tracks which we couldn't find")
	}

	albums, err := data.MusicRepository.SearchAlbum(track.ServiceArtistName, track.ServiceAlbumName)
	if err != nil {
		data.Logger.With(zap.Error(err)).Errorw("при поиске альбома в MasterDataRepository произошла ошибка",
			"Artist", track.ServiceArtistName,
			"Album", track.ServiceAlbumName)
	} else if len(albums) == 0 {
		data.Logger.Errorw("при поиске альбома в MasterDataRepository ничего не удалось найти",
			"Artist", track.ServiceArtistName,
			"Album", track.ServiceAlbumName)
	}

	// TODO: skip if don't have album
	masterAlbum := chooseAlbumFromSearchResults(albums, track.ServiceArtistName, track.ServiceAlbumName)
	if masterAlbum == nil {

		_, err = data.Engine.DataRepository.MissingReleaseRepository.Store(&core.MissingRelease{
			ArtistName: track.ServiceArtistName,
			AlbumName:  track.ServiceAlbumName,
		})

		if err != nil {
			data.Logger.With(zap.Error(err)).Error("при сохранении MissingRelease произошла ошибка")
		}

		return errors.New("couldn't find album")
	}

	track.ServiceArtistId = masterAlbum.ArtistMasterId
	track.ServiceAlbumId = masterAlbum.AlbumId

	return nil
}
