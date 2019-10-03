package core

import "time"

type Release struct {
	Id         uint
	AlbumId    uint
	PlaylistId uint
	SyncDate   time.Time

	AlbumName    string
	ArtistName   string
	Genres       string
	AlbumType    string
	ReleaseDate  string
	PlaylistName string
}
