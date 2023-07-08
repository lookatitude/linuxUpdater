package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewBrewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "brew",
		Short: "Manage brew packages",
		Run:   runBrewCmd,
	}

	return cmd
}

func runBrewCmd(cmd *cobra.Command, args []string) {
	brew := viper.Sub("brew")

	if checkBrewInstalled() {
		fmt.Println("Brew version:", getBrewVersion())
	} else {
		fmt.Println("No brew found")
		if brew.GetBool("install") {
			fmt.Println("Installing Homebrew...")
			err := exec.Command("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)").Run()
			if err != nil {
				fmt.Println("Failed to install Homebrew:", err)
				return
			}
		}
	}

	// Update brew and brew formulae. Brew itself is also updated in this step.
	if brew.GetBool("update") {
		fmt.Println("Running brew update...")
		err := exec.Command("brew", "update").Run()
		if err != nil {
			fmt.Println("Failed to update Homebrew:", err)
			return
		}
	}

	if brew.GetBool("upgrade") {
		fmt.Println("Running brew upgrade...")
		err := exec.Command("brew", "upgrade").Run()
		if err != nil {
			fmt.Println("Failed to upgrade Homebrew packages:", err)
			return
		}
	}

	if brew.GetBool("cleanup") {
		fmt.Println("Running brew cleanup...")
		err := exec.Command("brew", "cleanup").Run()
		if err != nil {
			fmt.Println("Failed to cleanup Homebrew packages:", err)
			return
		}
	}

	if brew.GetBool("doctor") {
		fmt.Println("Running brew doctor...")
		err := exec.Command("brew", "doctor").Run()
		if err != nil {
			fmt.Println("Failed to run brew doctor:", err)
			return
		}
	}
}

func checkBrewInstalled() bool {
	_, err := exec.LookPath("brew")
	return err == nil
}

func getBrewVersion() string {
	cmd := exec.Command("brew", "-v")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return ""
	}
	// The output will be something like "Homebrew 3.2.14\nHomebrew/homebrew-core (git revision 831ab; last commit 2021-09-27)\n"
	// We only want the first line.
	return strings.Split(out.String(), "\n")[0]
}
