package main

import (
	"os"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "tofu"
	app.Usage = "The cruelty-free elephant alternative."
	app.Commands = []*cli.Command{
		ClientCommand(),
		BlockCommand(),
		ServerCommand(),
		VolumeCommand(),
	}
	app.Run(os.Args)
}
