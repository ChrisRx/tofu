package main

import (
	"log"

	"gopkg.in/urfave/cli.v2"

	"github.com/ChrisRx/tofu/block"
)

func BlockCommand() *cli.Command {
	return &cli.Command{
		Name:  "block",
		Usage: "",
		Flags: []cli.Flag{},
		Subcommands: []*cli.Command{
			BlockGetCommand(),
			BlockListCommand(),
			BlockPutCommand(),
			BlockRunCommand(),
		},
	}
}

func BlockGetCommand() *cli.Command {
	return &cli.Command{
		Name:  "get",
		Usage: "",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			t := block.NewClient()
			defer t.Close()
			data, err := t.GetBlock(c.Args().First())
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Block data: %s\n", data)
			return nil
		},
	}
}

func BlockListCommand() *cli.Command {
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

func BlockPutCommand() *cli.Command {
	return &cli.Command{
		Name:  "put",
		Usage: "",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			t := block.NewClient()
			defer t.Close()
			b, err := t.PutBlock([]byte("lol"))
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Block: %s\n", b.Hash)
			return nil
		},
	}
}

func BlockRunCommand() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			s := block.NewBlockStore()
			s.Run()
			return nil
		},
	}
}
