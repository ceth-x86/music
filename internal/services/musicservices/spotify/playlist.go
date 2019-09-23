package spotify

import (
	"github.com/demas/music/internal/models/core"
	"github.com/zmb3/spotify"
	"go.uber.org/zap"
)

func (s *Spotify) DownloadPlaylist(playlistId string) (*core.Playlist, []*core.Track, error) {

	logger := zap.NewExample().Sugar()
	defer func() {
		_ = logger.Sync()
	}()

	spotifyPlaylist, err := s.Client.GetPlaylist(spotify.ID(playlistId))
	if err != nil {
		logger.With(zap.Error(err)).Error("error getting playlist from spotify")
		return nil, nil, err
	}

	playlist := convertPlaylist(spotifyPlaylist)
	playlist.PlaylistId = playlistId

	var tracks []*core.Track
	var total int
	offset := 0
	limit := 50

	for {
		spotifyTracks, err := s.Client.GetPlaylistTracksOpt(spotify.ID(playlistId), s.options(offset, limit), "")
		if err != nil {
			logger.With(zap.Error(err)).Error("error getting playlist tracks from spotify")
			return playlist, nil, err
		}
		total = spotifyTracks.Total

		for _, spotifyTrack := range spotifyTracks.Tracks {
			// sometimes we have such tracks from Spotify
			if spotifyTrack.Track.ID == "" {
				continue
			}
			tracks = append(tracks, convertTrack(&spotifyTrack))
		}

		offset += limit
		if total < offset {
			break
		}
	}

	return playlist, tracks, nil
}
