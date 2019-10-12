package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"sync"

	"github.com/demas/music/internal/models/enums"

	engine2 "github.com/demas/music/internal/services/engine"

	repository2 "github.com/demas/music/internal/services/datastore/repository"

	"github.com/alexeyco/simpletable"

	settings2 "github.com/demas/music/internal/services/settings"

	"github.com/demas/music/internal/models/core"
	"github.com/demas/music/internal/services/datastore"
	"github.com/demas/music/internal/services/datastore/dbutils"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var Service string
var Id string
var Filename string

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

		musicService, err := enums.ParseMusicService(Service)
		if err != nil {
			logger.With(zap.Error(err)).Fatal("Не удалось определить музыкальный сервис")
		}

		repository := datastore.NewPlaylistRepository(db)
		_, err = repository.Store(&core.Playlist{
			Service:    uint(musicService),
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

		table := simpletable.New()
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "Id"},
				{Align: simpletable.AlignCenter, Text: "Service"},
				{Align: simpletable.AlignCenter, Text: "PlaylistId"},
				{Align: simpletable.AlignCenter, Text: "Name"},
				{Align: simpletable.AlignCenter, Text: "Description"},
			}}

		for _, playlist := range repository.Fetch() {
			r := []*simpletable.Cell{
				{Text: fmt.Sprintf("%d", playlist.Id)},
				{Align: simpletable.AlignCenter, Text: enums.MusicService(playlist.Service).String()},
				{Text: playlist.PlaylistId},
				{Text: playlist.Name},
				{Text: playlist.Description},
			}
			table.Body.Cells = append(table.Body.Cells, r)
		}
		table.SetStyle(simpletable.StyleCompactLite)
		fmt.Println(table.String())
	},
}

var exportPlaylistCommand = &cobra.Command{
	Use:   "export",
	Short: "Export playlists",
	Long: `Examples:
		playlist export --filename playlists.json
	`,
	Run: func(cmd *cobra.Command, args []string) {

		logger := zap.NewExample().Sugar()
		defer func() {
			_ = logger.Sync()
		}()

		logger.Infow("ezport playlists")

		settings := settings2.InitSettings()
		db, err := dbutils.OpenDbConnection(settings.DbConnectionString, settings.TraceSqlCommand)
		if err != nil {
			logger.With(zap.Error(err)).Error("не удалось установить соединение с PostgreSQL")
		}

		repository := datastore.NewPlaylistRepository(db)
		playlists := repository.Fetch()

		data, err := json.Marshal(playlists)
		if err != nil {
			logger.With(zap.Error(err)).Error("Error playlists to JSON serialization")
			return
		}

		err = ioutil.WriteFile(Filename, data, 0644)
		if err != nil {
			logger.With(zap.Error(err)).Error("Error saving playlist to file")
		}

		fmt.Println("done")
	},
}

var syncPlaylistCommand = &cobra.Command{
	Use:   "sync",
	Short: "Sync playlists",
	Long: `Examples:
        playlist sync
		playlist sync --id 2
	`,
	Run: func(cmd *cobra.Command, args []string) {

		logger := zap.NewExample().Sugar()
		defer func() {
			_ = logger.Sync()
		}()

		logger.Infow("Sync playlist")

		settings := settings2.InitSettings()
		db, err := dbutils.OpenDbConnection(settings.DbConnectionString, settings.TraceSqlCommand)
		if err != nil {
			logger.With(zap.Error(err)).Error("не удалось установить соединение с PostgreSQL")
		}

		repository := repository2.NewRepository(db)
		engine := engine2.Engine{DataRepository: repository}
		var albums uint = 0
		var singles uint = 0
		var playlists uint = 0

		if Id != "" {

			logger.Infow("  Id = ", Id)
			idUint, err := strconv.ParseUint(Id, 10, 32)
			if err != nil {
				logger.With(zap.Error(err)).Errorw("Не удалось преобразовать идентификатор плейлиста в число",
					"Id", Id)
			}

			result := engine.PlaylistDownloader().Download(uint(idUint))
			albums = result.Album
			singles = result.Single

		} else {

			logger.Info("  all playlists")

			ch := make(chan *engine2.DownloadResult)
			var wg sync.WaitGroup

			increment := func(data *engine2.DownloadResult) {
				albums += data.Album
				singles += data.Single
				playlists += 1
			}

			go func() {
				for result := range ch {
					increment(result)
				}
			}()

			for _, playlist := range repository.PlaylistRepository.Fetch() {
				id := playlist.Id
				wg.Add(1)
				go func() {
					defer wg.Done()
					ch <- engine.PlaylistDownloader().Download(id)
				}()
			}

			wg.Wait()
			close(ch)
		}
		fmt.Printf("Обработано %d плейлистов\n", playlists)
		fmt.Printf("Загружено %d альбомов и %d синглов", albums, singles)
	},
}

func init() {
	rootCmd.AddCommand(playistCmd)

	playistCmd.AddCommand(addPlaylistCommand)
	addPlaylistCommand.Flags().StringVarP(&Service, "service", "s", "", "Music service")
	addPlaylistCommand.Flags().StringVarP(&Id, "id", "i", "", "Playlist id")
	_ = addPlaylistCommand.MarkFlagRequired("service")
	_ = addPlaylistCommand.MarkFlagRequired("id")

	playistCmd.AddCommand(syncPlaylistCommand)
	syncPlaylistCommand.Flags().StringVarP(&Id, "id", "i", "", "Playlist id")

	playistCmd.AddCommand(exportPlaylistCommand)
	exportPlaylistCommand.Flags().StringVarP(&Filename, "filename", "f", "playlist.json", "Filename")
	_ = exportPlaylistCommand.MarkFlagRequired("filename")

	playistCmd.AddCommand(listPlaylistCommand)
}
