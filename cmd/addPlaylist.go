package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var Service string
var Id string

// addPlaylistCmd represents the addPlaylist command
var addPlaylistCmd = &cobra.Command{
	Use:   "addPlaylist",
	Short: "Add new playlist",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		sugar := zap.NewExample().Sugar()
		defer sugar.Sync()
		sugar.Infow("adding playlist",
			"service", Service,
			"id", Id)
	},
}

func init() {
	rootCmd.AddCommand(addPlaylistCmd)
	addPlaylistCmd.Flags().StringVarP(&Service, "service", "s", "", "Music service")
	addPlaylistCmd.Flags().StringVarP(&Id, "id", "i", "", "Playlist id")
}
