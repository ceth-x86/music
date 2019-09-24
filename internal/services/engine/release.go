package engine

import "time"

const daysForRelease = 30

func isItNewRelease(releaseDate time.Time) bool {
	if time.Now().Sub(releaseDate).Hours() < 24*daysForRelease {
		return true
	}

	return false
}
