package settings

import "github.com/namsral/flag"

type Settings struct {
	DbConnectionString string
	TraceSqlCommand    bool
}

func InitSettings() *Settings {
	settings := new(Settings)
	flag.StringVar(&settings.DbConnectionString, "dbConnectionString", "none", "Connection string to PostgreSQL")
	flag.BoolVar(&settings.TraceSqlCommand, "traceSqlCommands", true, "Trace sql commands")
	flag.Parse()
	return settings
}
