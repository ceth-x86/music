package cmd

import (
	"fmt"

	settings2 "github.com/demas/music/internal/services/settings"

	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/services/datastore"
	"github.com/demas/music/internal/services/datastore/dbutils"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var Service string
var Id string

// addPlaylistCmd represents the addPlaylist command
var playistCmd = &cobra.Command{
	Use:   "playlist",
	Short: "Working with playlists",
	Args:  cobra.MinimumNArgs(1),
}

var addPlaylistCommand = &cobra.Command{
	Use:   "add",
	Short: "Adding playlists",
	Long: `Adding playlists. Examples:
		playlist add --service spoify --id dAKJ23K4H234
		playlist add --service deezer --id 23424
		playlist add --service yandex --id 3123123123
	`,
	Run: func(cmd *cobra.Command, args []string) {

		logger := zap.NewExample().Sugar()
		defer func() {
			_ = logger.Sync()
		}()

		logger.Infow("adding playlist",
			"service", Service,
			"id", Id)

		settings := settings2.InitSettings()
		db, err := dbutils.OpenDbConnection(settings.DbConnectionString, settings.TraceSqlCommand)
		if err != nil {
			logger.With(zap.Error(err)).Error("не удалось установить соединение с PostgreSQL")
		}

		repository := datastore.NewPlaylistRepository(db)
		_, err = repository.Store(&core.Playlist{
			Service:    Service,
			PlaylistId: Id,
		})
		if err != nil {
			logger.With(zap.Error(err)).Error("не удалось создать плейлист")
		}
	},
}

var listPlaylistCommand = &cobra.Command{
	Use:   "list",
	Short: "Show playlists",
	Long: `Examples:
		playlist show
	`,
	Run: func(cmd *cobra.Command, args []string) {

		logger := zap.NewExample().Sugar()
		defer func() {
			_ = logger.Sync()
		}()

		logger.Infow("show list of playlists")

		settings := settings2.InitSettings()
		db, err := dbutils.OpenDbConnection(settings.DbConnectionString, settings.TraceSqlCommand)
		if err != nil {
			logger.With(zap.Error(err)).Error("не удалось установить соединение с PostgreSQL")
		}

		repository := datastore.NewPlaylistRepository(db)
		for _, playlist := range repository.Fetch() {
			fmt.Println(playlist.Service, playlist.Id)
		}

	},
}

func init() {
	rootCmd.AddCommand(playistCmd)

	playistCmd.AddCommand(addPlaylistCommand)
	addPlaylistCommand.Flags().StringVarP(&Service, "service", "s", "", "Music service")
	addPlaylistCommand.Flags().StringVarP(&Id, "id", "i", "", "Playlist id")
	_ = addPlaylistCommand.MarkFlagRequired("service")
	_ = addPlaylistCommand.MarkFlagRequired("id")

	playistCmd.AddCommand(listPlaylistCommand)
}
