package yandex

import (
	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/models/enums"
	"github.com/demas/yandexmusic"
)

// TODO: remove demas/music.models dependency

type Yandex struct{}

func NewYandex() *Yandex {
	return &Yandex{}
}

func convertPlaylist(yandexPlaylist *yandexmusic.MusicPlaylist) *core.Playlist {
	return &core.Playlist{
		Service:     uint(enums.MusicServiceYandex),
		Name:        yandexPlaylist.Name,
		Description: "",
	}
}

func convertTrack(yandexTrack *yandexmusic.MusicPlaylistTrack) *core.Track {
	return &core.Track{
		// PlaylistId: 0,
		Name:              yandexTrack.Name,
		TrackId:           yandexTrack.YandexId,
		ServiceAlbumName:  yandexTrack.Album,
		ServiceArtistName: yandexTrack.Artist,
		MasterData:        false,
	}
}
