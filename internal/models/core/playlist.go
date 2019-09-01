package core

import (
	"time"
)

type Playlist struct {
	Id          uint
	Service     string
	PlaylistId  string
	TrackCount  *uint
	LastChanged *time.Time
}
