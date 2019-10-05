package core

import "time"

type Album struct {
	Id          uint
	Name        string
	AlbumId     string
	ArtistId    uint
	ReleaseDate time.Time
	Popularity  int
	// TODO: may be enum
	AlbumType string

	ArtistMasterId    string
	ArtistName        string
	ReleaseDateString string
}
