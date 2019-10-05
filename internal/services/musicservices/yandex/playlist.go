package yandex

import (
	"strconv"

	"github.com/demas/music/internal/models/core"
	"github.com/demas/yandexmusic"
	"go.uber.org/zap"
)

func (y *Yandex) DownloadPlaylist(playlistId string) (*core.Playlist, []*core.Track, error) {

	logger := zap.NewExample().Sugar()
	defer func() {
		_ = logger.Sync()
	}()

	idInt, err := strconv.Atoi(playlistId)
	if err != nil {
		logger.With(zap.Error(err)).Errorw("cannot convert playlistId to int for Yandex playlist",
			"playlistId", playlistId)
	}

	yandexPlaylist := yandexmusic.GetPlaylist(int64(idInt))
	playlist := convertPlaylist(yandexPlaylist)
	playlist.PlaylistId = playlistId

	var tracks []*core.Track
	for _, yandexTrack := range yandexPlaylist.Tracks {
		tracks = append(tracks, convertTrack(yandexTrack))

	}

	return playlist, tracks, nil
}
