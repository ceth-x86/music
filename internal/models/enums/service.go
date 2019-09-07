//go:generate go-enum -f=$GOFILE --marshal
package enums

// MusicService is an enumeration of music services that are allowed
/*
ENUM(
spotify
deezer
yandex
)
*/
type MusicService uint
