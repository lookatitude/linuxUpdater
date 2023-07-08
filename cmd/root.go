package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "linux-updater",
	Short: "linux-updater is a CLI tool for updating Linux system and managing packages",
	Long:  `linux-updater is a CLI tool for updating Linux system and managing packages. It supports updating system packages using apt and brew.`,
	Run: func(cmd *cobra.Command, args []string) {
		aptCmd := NewAptCmd()
		aptCmd.Run(aptCmd, args)
		brewCmd := NewBrewCmd()
		brewCmd.Run(brewCmd, args)
		snapCmd := NewSnapCmd()
		snapCmd.Run(snapCmd, args)
		flatpakCmd := NewFlatpakCmd()
		flatpakCmd.Run(flatpakCmd, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Add commands
	RootCmd.AddCommand(NewAptCmd())
	RootCmd.AddCommand(NewBrewCmd())
	RootCmd.AddCommand(NewSnapCmd())
	RootCmd.AddCommand(NewFlatpakCmd())
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.linux-updater")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
