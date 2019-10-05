package core

type Track struct {
	Id         uint
	PlaylistId uint
	Name       string
	TrackId    string
	ArtistId   uint
	AlbumId    uint

	MasterData        bool
	ServiceArtistId   string
	ServiceAlbumId    string
	ServiceArtistName string
	ServiceAlbumName  string
}
