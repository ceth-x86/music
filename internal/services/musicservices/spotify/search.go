package spotify

import (
	"errors"

	"github.com/demas/music/internal/models/core"

	"github.com/zmb3/spotify"
)

func (s *Spotify) SearchAlbum(artist string, album string) ([]*core.Album, error) {

	searchResult, err := s.Client.Search(artist+" "+album, spotify.SearchTypeAlbum)
	if err != nil {
		return nil, err
	}

	if searchResult == nil {
		return nil, errors.New("searchResult is nil")
	}

	if searchResult.Albums == nil {
		return nil, errors.New("searchResult.Albums is nil")
	}

	var result []*core.Album
	for _, spotifyAlbum := range searchResult.Albums.Albums {
		result = append(result, convertAlbum(&spotifyAlbum))
	}

	return result, nil
}
