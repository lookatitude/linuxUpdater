package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewAptCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apt",
		Short: "Manage apt packages",
		Run:   runAptCmd,
	}

	return cmd
}

func runAptCmd(cmd *cobra.Command, args []string) {
	apt := viper.Sub("apt")

	if apt.GetBool("update") {
		fmt.Println("Running apt-get update...")
		err := exec.Command("sudo", "apt-get", "update", "-y").Run()
		if err != nil {
			fmt.Println("Failed to update apt packages:", err)
			return
		}
	}

	if apt.GetBool("upgrade") {
		fmt.Println("Running apt-get upgrade...")
		err := exec.Command("sudo", "apt-get", "upgrade", "-y").Run()
		if err != nil {
			fmt.Println("Failed to upgrade apt packages:", err)
			return
		}
	}

	if apt.GetBool("distUpgrade") {
		fmt.Println("Running apt-get dist-upgrade...")
		err := exec.Command("sudo", "apt-get", "dist-upgrade", "-y").Run()
		if err != nil {
			fmt.Println("Failed to dist-upgrade apt packages:", err)
			return
		}
	}

	if apt.GetBool("autoremove") {
		fmt.Println("Running apt-get autoremove...")
		err := exec.Command("sudo", "apt-get", "autoremove", "-y").Run()
		if err != nil {
			fmt.Println("Failed to autoremove apt packages:", err)
			return
		}
	}

	if apt.GetBool("autoclean") {
		fmt.Println("Running apt-get autoclean...")
		err := exec.Command("sudo", "apt-get", "autoclean", "-y").Run()
		if err != nil {
			fmt.Println("Failed to autoclean apt packages:", err)
			return
		}
	}
}
