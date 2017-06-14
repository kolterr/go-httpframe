package command

import (
	"go-httpframe/internal/app"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "lilin",
	Short: "service",
	Long:  "author is lilin",
	Run: func(cmd *cobra.Command, args []string) {
		app.Run(NewServerApp("lilin", cmd.Flags().Lookup("conf").Value.String()))
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.Flags().StringP("conf", "c", "configs/lilin.toml", "the path to the config file")
}
