package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "calorie-api",
	Short: "The API tier for the Calorie Tracker API",
	Long:  `The API tier for the Calorie Tracker API, which allows for saving of calories consumed during the day.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var cfgFile string

func init() {
	cobra.OnInitialize(InitConfig)

	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is ./config/local-config.yml")
}

func InitConfig() {
	viper.SetEnvPrefix("calories_api")
	viper.AutomaticEnv()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./config/")
		viper.SetConfigType("yml")
		viper.SetConfigName("local-config")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't read config: ", err)
	}

	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(getConfiguredLogLevel())

	log.Debugf("Running with config:")
	keys := viper.AllKeys()
	for _, key := range keys {
		log.Debugf("    %s : %v", key, viper.Get(key))
	}
}

func getConfiguredLogLevel() log.Level {
	logLevel := viper.GetString("logLevel")
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("Invalid log level supplied")
	}
	return level
}
