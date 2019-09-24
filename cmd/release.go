package cmd

import (
	"fmt"

	"github.com/demas/music/internal/models/core"

	"github.com/alexeyco/simpletable"

	settings2 "github.com/demas/music/internal/services/settings"

	"github.com/demas/music/internal/services/datastore"
	"github.com/demas/music/internal/services/datastore/dbutils"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var AlbumType string

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Working with releases",
}

var listReleaseCommand = &cobra.Command{
	Use:   "list",
	Short: "Show releases",
	Long: `Examples:
		release list
	`,
	Run: func(cmd *cobra.Command, args []string) {

		logger := zap.NewExample().Sugar()
		defer func() {
			_ = logger.Sync()
		}()

		logger.Infow("show list of releases")

		settings := settings2.InitSettings()
		db, err := dbutils.OpenDbConnection(settings.DbConnectionString, settings.TraceSqlCommand)
		if err != nil {
			logger.With(zap.Error(err)).Error("не удалось установить соединение с PostgreSQL")
		}

		repository := datastore.NewReleaseRepository(db)

		var releases []*core.Release
		if AlbumType == "" {
			releases = repository.Fetch()
		} else if AlbumType == "album" {
			releases = repository.GetByAlbumType(AlbumType)
		} else {
			fmt.Println("unknown album type")
			return
		}

		table := simpletable.New()
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "Id"},
				{Align: simpletable.AlignCenter, Text: "Artist"},
				{Align: simpletable.AlignCenter, Text: "Album"},
				{Align: simpletable.AlignCenter, Text: "Type"},
				{Align: simpletable.AlignCenter, Text: "Genres"},
			}}

		for _, release := range releases {
			r := []*simpletable.Cell{
				{Text: fmt.Sprintf("%d", release.Id)},
				{Text: release.ArtistName},
				{Text: release.AlbumName},
				{Text: release.AlbumType},
				{Text: release.Genres},
			}
			table.Body.Cells = append(table.Body.Cells, r)
		}
		table.SetStyle(simpletable.StyleCompactLite)
		fmt.Println(table.String())
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)
	releaseCmd.AddCommand(listReleaseCommand)
	listReleaseCommand.Flags().StringVarP(&AlbumType, "type", "t", "", "Release type (album)")
}
