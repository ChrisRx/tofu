package main

import (
	"log"
	"path/filepath"

	"github.com/coreos/etcd/embed"
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
			_, err := startEtcd()
			if err != nil {
				log.Fatal(err)
			}
			v := volume.New()
			v.Run()
			return nil
		},
	}
}

func startEtcd() (*embed.Etcd, error) {
	c := embed.NewConfig()
	c.Name = "default"
	c.Dir = filepath.Join("data", "etcd")
	c.WalDir = ""
	return embed.StartEtcd(c)
}
