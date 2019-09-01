package cmd

import (
	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/services/datastore"
	"github.com/demas/music/internal/services/datastore/dbutils"
	appSettings "github.com/demas/music/internal/services/settings"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var Service string
var Id string

// addPlaylistCmd represents the addPlaylist command
var addPlaylistCmd = &cobra.Command{
	Use:   "addPlaylist",
	Short: "Add new playlist",
	Long: `Add new playlist. Examples:
		addPlaylist --service spoify --id dAKJ23K4H234
		addPlaylist --service deezer --id 23424
		addPlaylist --service yandex --id 3123123123
	`,
	Run: func(cmd *cobra.Command, args []string) {

		logger := zap.NewExample().Sugar()
		defer func() {
			_ = logger.Sync()
		}()

		logger.Infow("adding playlist",
			"service", Service,
			"id", Id)

		settings := appSettings.InitSettings()
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

func init() {
	rootCmd.AddCommand(addPlaylistCmd)
	addPlaylistCmd.Flags().StringVarP(&Service, "service", "s", "", "Music service")
	addPlaylistCmd.Flags().StringVarP(&Id, "id", "i", "", "Playlist id")
}
