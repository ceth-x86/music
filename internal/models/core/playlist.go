package core

import (
	"time"
)

type Playlist struct {
	Id          uint
	Service     uint
	PlaylistId  string
	Name        string
	Description string
	TrackCount  *uint
	LastChanged *time.Time
}
