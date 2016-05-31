package main

import (
	//"log"

	"gopkg.in/urfave/cli.v2"

	"github.com/ChrisRx/tofu/volume"
)

func RunCommand() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			s := volume.New()
			s.Run()
			return nil
		},
	}
}

func VolumeCommand() *cli.Command {
	return &cli.Command{
		Name:  "volume",
		Usage: "",
		Flags: []cli.Flag{},
		Subcommands: []*cli.Command{
			RunCommand(),
		},
	}
}
