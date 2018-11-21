package main

import (
	"os"

	"github.com/ObsidianRock/silsila/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
