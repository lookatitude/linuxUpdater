package main

import (
	"github.com/lookatitude/linuxUpdater/cmd"
	"github.com/spf13/cobra/doc"
	"log"
)

func main() {

	// Generate command documentation
	err := doc.GenMarkdownTree(cmd.RootCmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}

	// Run the application
	cmd.Execute()
}
