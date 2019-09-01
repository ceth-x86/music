package cmd

import (
	"os"

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
var addPlaylistCmd = &cobra.Command{
	Use:   "playlist",
	Short: "Working with playlists",
	Long: `Working with playlists. Examples:
		playlist add --service spoify --id dAKJ23K4H234
		playlist add --service deezer --id 23424
		playlist add --service yandex --id 3123123123
	`,
	Run: func(cmd *cobra.Command, args []string) {

		logger := zap.NewExample().Sugar()
		defer func() {
			_ = logger.Sync()
		}()

		if len(args) == 0 {
			logger.Errorw("command is missing")
			os.Exit(0)
		}

		switch args[0] {
		case "add":
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
		default:
			logger.Errorw("unknown command")
		}
	},
}

func init() {
	rootCmd.AddCommand(addPlaylistCmd)
	addPlaylistCmd.Flags().StringVarP(&Service, "service", "s", "", "Music service")
	addPlaylistCmd.Flags().StringVarP(&Id, "id", "i", "", "Playlist id")
}
