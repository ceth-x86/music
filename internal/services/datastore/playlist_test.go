package datastore

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func TestGetPlaylistById(t *testing.T) {

	var (
		id          = 1
		service     = 1
		playlist_id = "123"
	)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(`^SELECT \* FROM "playlists*"`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "service", "playlist_id"}).
			AddRow(id, service, playlist_id))

	gormDB, _ := gorm.Open("postgres", db)
	repository := NewPlaylistRepository(gormDB.LogMode(true))
	_, err = repository.GetById(uint(id))

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
