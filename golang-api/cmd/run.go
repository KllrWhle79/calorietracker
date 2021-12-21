package cmd

import (
	"github.com/KllrWhle79/calorietracker/api"
	"github.com/KllrWhle79/calorietracker/db"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the API",
	Long:  `Run the DR Client API service.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := db.InitTables(false)
		if err != nil {
			panic(err)
		}
		api.Start()
	},
}

func init() {
	RootCmd.AddCommand(RunCmd)
}
