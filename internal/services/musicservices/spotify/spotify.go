package spotify

import (
	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/models/enums"
	"github.com/zmb3/spotify"
	"go.uber.org/zap"
)

type Spotify struct {
	Client spotify.Client
}

func NewSpotify() *Spotify {
	return &Spotify{Client: createClient()}
}

func (s *Spotify) options(offset int, limit int) *spotify.Options {
	return &spotify.Options{
		Offset: &offset,
		Limit:  &limit,
	}
}

func convertTrack(spotifyTrack *spotify.PlaylistTrack) *core.Track {
	return &core.Track{
		// PlaylistId: 0,
		Name:    spotifyTrack.Track.Name,
		TrackId: string(spotifyTrack.Track.ID),
	}
}

func (s *Spotify) DownloadPlaylist(playlistId string) (*core.Playlist, []*core.Track, error) {

	logger := zap.NewExample().Sugar()
	//defer func() {
	//	_ = logger.Sync()
	//}()

	// понадобится позже для заполнения полей Плейлиста
	spotifyPlaylist, err := s.Client.GetPlaylist(spotify.ID(playlistId))
	if err != nil {
		logger.With(zap.Error(err)).Error("error getting playlist from spotify")
		return nil, nil, err
	}

	playlist := &core.Playlist{
		Service:     uint(enums.MusicServiceSpotify),
		PlaylistId:  playlistId,
		Name:        spotifyPlaylist.Name,
		Description: spotifyPlaylist.Description,
	}

	var tracks []*core.Track
	var total int
	offset := 0
	limit := 50

	for {
		spotifyTracks, err := s.Client.GetPlaylistTracksOpt(spotify.ID(playlistId), s.options(offset, limit), "")
		if err != nil {
			return playlist, nil, err
		}
		total = spotifyTracks.Total

		for _, spotifyTrack := range spotifyTracks.Tracks {
			tracks = append(tracks, convertTrack(&spotifyTrack))
		}

		offset += limit
		if total < offset {
			break
		}
	}

	return playlist, tracks, nil
}
