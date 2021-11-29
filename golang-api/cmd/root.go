package cmd

import (
	"github.com/KllrWhle79/calorietracker/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "calorie-api",
	Short: "The API tier for the Calorie Tracker API",
	Long:  `The API tier for the Calorie Tracker API, which allows for saving of calories consumed during the day.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	cobra.OnInitialize(config.InitConfig)
}
