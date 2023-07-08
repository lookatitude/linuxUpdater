package cmd

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewFlatpakCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "flatpak",
		Short: "Manage flatpak packages",
		Run:   runFlatpakCmd,
	}

	cmd.Flags().BoolP("update", "u", false, "Update flatpak packages")

	return cmd
}

func runFlatpakCmd(cmd *cobra.Command, args []string) {
	flatpak := viper.Sub("flatpak")

	if checkFlatpakInstalled() {
		fmt.Println("Flatpak version:", getFlatpakVersion())
		if flatpak.GetBool("update") {
			fmt.Println("Updating flatpak packages...")
			err := exec.Command("flatpak", "update").Run()
			if err != nil {
				fmt.Println("Failed to update flatpak packages:", err)
				return
			}
		}
	} else {
		fmt.Println("No flatpak found")
	}
}

func checkFlatpakInstalled() bool {
	_, err := exec.LookPath("flatpak")
	return err == nil
}

func getFlatpakVersion() string {
	cmd := exec.Command("flatpak", "--version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return ""
	}
	return out.String()
}
