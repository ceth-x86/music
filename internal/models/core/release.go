package core

import "time"

type Release struct {
	Id       uint
	AlbumId  uint
	SyncDate time.Time
}
