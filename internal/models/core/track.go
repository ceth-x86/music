package core

// created during playlist synchronization process (`playlist sync`)
type Track struct {
	Id         uint
	PlaylistId uint   // playlist ID in the database
	Name       string // track name (music service)
	TrackId    string // track ID (music service)
	ArtistId   uint   // artist ID (database)
	AlbumId    uint   // album ID (database)

	// calculated columns
	MasterData        bool // true if the track belongs to Spotify playlist (MasterData)
	ServiceArtistId   string
	ServiceAlbumId    string
	ServiceArtistName string
	ServiceAlbumName  string
}
