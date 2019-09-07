package core

import (
	"time"
)

type Playlist struct {
	Id          uint
	Service     uint
	PlaylistId  string
	TrackCount  *uint
	LastChanged *time.Time
}
