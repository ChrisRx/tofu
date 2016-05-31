package main

import (
	//"log"

	"gopkg.in/urfave/cli.v2"

	"github.com/ChrisRx/tofu/block"
	"github.com/ChrisRx/tofu/volume"
)

func ServerCommand() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			s := block.NewBlockStore()
			go s.Run()
			v := volume.New()
			v.Run()
			return nil
		},
	}
}
