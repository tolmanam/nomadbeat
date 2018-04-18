package main

import (
	"os"

	"github.com/tolmanam/nomadbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
