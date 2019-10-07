package spotify

import (
	"sync"

	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/models/enums"
	"github.com/zmb3/spotify"
)

var (
	once     sync.Once
	instance *Spotify
)

type Spotify struct {
	Client spotify.Client
}

func NewSpotify() *Spotify {
	once.Do(func() {
		instance = &Spotify{Client: createClient()}
	})
	return instance
}

func (s *Spotify) options(offset int, limit int) *spotify.Options {
	return &spotify.Options{
		Offset: &offset,
		Limit:  &limit,
	}
}

func convertPlaylist(spotifyPlaylist *spotify.FullPlaylist) *core.Playlist {
	return &core.Playlist{
		Service:     uint(enums.MusicServiceSpotify),
		Name:        spotifyPlaylist.Name,
		Description: spotifyPlaylist.Description,
	}
}

func convertTrack(spotifyTrack *spotify.PlaylistTrack) *core.Track {
	return &core.Track{
		// PlaylistId: 0,
		Name:            spotifyTrack.Track.Name,
		TrackId:         string(spotifyTrack.Track.ID),
		ServiceAlbumId:  string(spotifyTrack.Track.Album.ID),
		ServiceArtistId: string(spotifyTrack.Track.Artists[0].ID),
		MasterData:      true,
	}
}

func convertAlbum(spotifyAlbum *spotify.SimpleAlbum) *core.Album {
	return &core.Album{
		Name:              spotifyAlbum.Name,
		AlbumId:           string(spotifyAlbum.ID),
		AlbumType:         spotifyAlbum.AlbumType,
		ArtistMasterId:    string(spotifyAlbum.Artists[0].ID),
		ArtistName:        spotifyAlbum.Artists[0].Name,
		ReleaseDateString: spotifyAlbum.ReleaseDate,
	}
}
