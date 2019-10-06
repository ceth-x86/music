package spotify

import (
	"fmt"

	"github.com/demas/music/internal/models/core"

	"github.com/zmb3/spotify"
)

func (s *Spotify) SearchAlbum(artist string, album string) []*core.Album {

	searchResult, err := s.Client.Search(artist+" "+album, spotify.SearchTypeAlbum)
	if err != nil {
		// TODO: logging
		fmt.Println(err.Error())
	}

	var result []*core.Album
	if searchResult != nil && searchResult.Albums != nil {
		for _, spotifyAlbum := range searchResult.Albums.Albums {
			result = append(result, convertAlbum(&spotifyAlbum))
		}
	}

	return result
}
