package core

import (
	"time"
)

// created manually using `playlist add` command
type Playlist struct {
	Id          uint
	Service     uint       // music service (enum - Spotify, Deezer, Yandex)
	PlaylistId  string     // playlist's ID in the music service
	Name        string     // playlist's name (filled from music service)
	Description string     // playlist's description (filled from music service)
	TrackCount  *uint      // TODO: not used
	LastChanged *time.Time // time of last change of the playlist
}
