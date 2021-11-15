package cmd

import (
	"github.com/KllrWhle79/calorietracker/api"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the API",
	Long:  `Run the DR Client API service.`,
	Run: func(cmd *cobra.Command, args []string) {
		api.Start()
	},
}

func init() {
	RootCmd.AddCommand(RunCmd)
}
