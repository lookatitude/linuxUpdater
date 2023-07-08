package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"log"
)

func NewDocCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "doc",
		Short: "Generate documentation",
		Run:   runDocCmd,
	}

	cmd.Flags().BoolP("markdown", "m", false, "Generate markdown documentation")

	return cmd
}

func runDocCmd(cmd *cobra.Command, args []string) {
	markdown, _ := cmd.Flags().GetBool("markdown")

	if markdown {
		err := doc.GenMarkdownTree(RootCmd, "./docs")
		if err != nil {
			log.Fatal(err)
		}
	}
}
