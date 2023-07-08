package cmd

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewSnapCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "snap",
		Short: "Manage snap packages",
		Run:   runSnapCmd,
	}

	cmd.Flags().BoolP("refresh", "r", false, "Refresh snap packages")

	return cmd
}

func runSnapCmd(cmd *cobra.Command, args []string) {
	snap := viper.Sub("snap")

	if checkSnapInstalled() {
		fmt.Println("Snap version:", getSnapVersion())
		if snap.GetBool("refresh") {
			fmt.Println("Refreshing snap packages...")
			err := exec.Command("sudo", "snap", "refresh").Run()
			if err != nil {
				fmt.Println("Failed to refresh snap packages:", err)
				return
			}
		}
	} else {
		fmt.Println("No snap found")
	}
}

func checkSnapInstalled() bool {
	_, err := exec.LookPath("snap")
	return err == nil
}

func getSnapVersion() string {
	cmd := exec.Command("snap", "--version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return ""
	}
	return out.String()
}
