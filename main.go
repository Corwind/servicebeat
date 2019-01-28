package main

import (
	"os"

	"github.com/Corwind/servicebeat/cmd"

	_ "github.com/Corwind/servicebeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
