package main

import (
	"os"
        "os/exec"

	"github.com/Corwind/servicebeat/cmd"

	_ "github.com/Corwind/servicebeat/include"
)

func main() {
        exec.Command("date").Run()
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
