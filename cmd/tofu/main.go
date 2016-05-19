package main

import (
	"log"
	"os"

	"tofu/client"
	"tofu/store/block"

	"gopkg.in/codegangsta/cli.v2"
)

func FileSystemCommand() {
	return cli.Command{
		Name:    "mount",
		Aliases: []string{"m"},
		Usage:   "mount TofuFS to sub-directory",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "readonly, r",
				Usage: "mount file system as read-only",
			},
		},
		Action: func(c *cli.Context) error {
			// Should check if dir exists or not.
			// Maybe check to ensure that only relative path is allowed also.
			// Should also try and fix a messed up mountpoint by attempting an unmount since
			// it is possible that the program doesn't call close when it exists abnormally
			tfs, err := NewTofuFS()
			if err != nil {
				log.Fatal(err)
			}
			if err = tfs.MountAndServe(c.Args().First(), c.Bool("readonly")); err != nil {
				log.Fatal(err)
			}
			return nil
		},
	}
}

func clientCommand() cli.Command {
	return cli.Command{
		Name:  "client",
		Usage: "",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			t := client.NewTofuClient()
			defer t.Close()
			for _, block := range t.ListBlocks() {
				log.Printf("Block %s\n", block)
			}
			return nil
		},
	}
}

func blockStoreCommand() cli.Command {
	return cli.Command{
		Name:  "block",
		Usage: "",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			s := block.NewBlockStore()
			s.Run()
			return nil
		},
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "tofu"
	app.Usage = "distributed file system"
	app.Commands = []cli.Command{
		clientCommand(),
		blockStoreCommand(),
	}
	app.Run(os.Args)
}
