package command

import (
	"github.com/spf13/cobra"
)

const version = "0.0.1"

//The version command prints this service.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version.",
	Long:  "The version of the lilin service.",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(logo)
		println("lilin version ", version)
	},
}
