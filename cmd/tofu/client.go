package main

import (
	"log"

	"gopkg.in/urfave/cli.v2"

	"github.com/ChrisRx/tofu/block"
	"github.com/ChrisRx/tofu/client"
)

func ClientCommand() *cli.Command {
	return &cli.Command{
		Name:  "client",
		Usage: "",
		Flags: []cli.Flag{},
		Subcommands: []*cli.Command{
			GetCommand(),
			ListCommand(),
			PutCommand(),
		},
	}
}

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:  "get",
		Usage: "",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			t := client.NewTofuClient()
			if c.NArg() == 0 {
				log.Fatal("Not enough args")
			}
			t.GetFile(c.Args().First())
			return nil
		},
	}
}

func ListCommand() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			t := block.NewClient()
			defer t.Close()
			for _, block := range t.ListBlocks() {
				log.Printf("Block %s\n", block)
			}
			return nil
		},
	}
}

func PutCommand() *cli.Command {
	return &cli.Command{
		Name:  "put",
		Usage: "",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			t := client.NewTofuClient()
			//defer t.Close()
			if c.NArg() == 0 {
				log.Fatal("Not enough args")
			}
			t.PutFile(c.Args().First())
			return nil
		},
	}
}
