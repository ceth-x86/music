package settings

import (
	"sync"

	"github.com/namsral/flag"
)

type Settings struct {
	DbConnectionString  string
	TraceSqlCommand     bool
	SpotifyClientId     string
	SpotifyClientSecret string
}

var settings *Settings
var once sync.Once

func InitSettings() *Settings {
	once.Do(func() {
		settings = new(Settings)
		flag.StringVar(&settings.DbConnectionString, "dbConnectionString", "none", "Connection string to PostgreSQL")
		flag.BoolVar(&settings.TraceSqlCommand, "traceSqlCommands", false, "Trace sql commands")
		flag.StringVar(&settings.SpotifyClientId, "spotifyClientId", "", "Spotify ClientID")
		flag.StringVar(&settings.SpotifyClientSecret, "spotifyClientSecret", "", "Spotify Client Secret")
		flag.Parse()
	})
	return settings
}
