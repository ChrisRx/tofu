package main

import (
	"os"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := cli.App{
		Name:  "tofu",
		Usage: "The cruelty-free elephant alternative.",
		Commands: []*cli.Command{
			ClientCommand(),
			BlockCommand(),
			ServerCommand(),
			VolumeCommand(),
		},
	}
	app.Run(os.Args)
}
