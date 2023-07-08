package main

import (
	"github.com/lookatitude/linuxUpdater/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "linux-updater",
	Short: "linux-updater is a CLI tool for updating Linux system and managing packages",
	Long:  `linux-updater is a CLI tool for updating Linux system and managing packages. It supports updating system packages using apt and brew.`,
}

func init() {
	cobra.OnInitialize(initConfig)

	// Add commands
	rootCmd.AddCommand(cmd.NewAptCmd())
	rootCmd.AddCommand(cmd.NewBrewCmd())
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
