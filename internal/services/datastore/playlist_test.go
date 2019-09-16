package datastore

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/demas/music/internal/models/core"

	"github.com/stretchr/testify/suite"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository *PlaylistRepository
	playlist   *core.Playlist
}

func (suite *Suite) SetupTest() {

	var (
		db  *sql.DB
		err error
	)

	db, suite.mock, err = sqlmock.New()
	require.NoError(suite.T(), err)

	suite.DB, err = gorm.Open("postgres", db)
	require.NoError(suite.T(), err)

	suite.DB.LogMode(true)
	suite.repository = NewPlaylistRepository(suite.DB)
}

func (suite *Suite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *Suite) TestFetch() {

	suite.mock.ExpectQuery(`^SELECT \* FROM "playlists" WHERE \(deleted_at is null\)`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "service", "playlist_id"}))

	suite.repository.Fetch()
}

func (suite *Suite) TestGetPlaylistById() {

	var (
		id          = 1
		service     = 1
		playlist_id = "123"
	)

	suite.mock.ExpectQuery(`^SELECT \* FROM "playlists" WHERE ("playlist"."id" = ?)*`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "service", "playlist_id"}).
			AddRow(id, service, playlist_id))

	_, err := suite.repository.GetById(uint(id))
	require.NoError(suite.T(), err)
}

func (suite *Suite) TestStore() {

	var (
		returnId   uint = 1
		service    uint = 1
		playlistId      = "123"
	)

	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(`^INSERT INTO "playlists"`).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(returnId))
	suite.mock.ExpectCommit()

	_, err := suite.repository.Store(&core.Playlist{Service: service, PlaylistId: playlistId})
	require.NoError(suite.T(), err)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
