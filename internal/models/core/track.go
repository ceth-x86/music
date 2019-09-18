package core

type Track struct {
	Id         uint
	PlaylistId uint
	Name       string
	TrackId    string
	ArtistId   uint
	AlbumId    uint

	ServiceArtistId string
	ServiceAlbumId  string
}
