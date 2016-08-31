package main

import (
	//"log"
	//"os"

	"gopkg.in/urfave/cli.v2"
)

func FileSystemCommand() *cli.Command {
	return &cli.Command{
		Name:    "mount",
		Aliases: []string{"m"},
		Usage:   "mount TofuFS to sub-directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "readonly, r",
				Usage: "mount file system as read-only",
			},
		},
		Action: func(c *cli.Context) error {
			// Should check if dir exists or not.
			// Maybe check to ensure that only relative path is allowed also.
			// Should also try and fix a messed up mountpoint by attempting an unmount since
			// it is possible that the program doesn't call close when it exists abnormally
			//tfs, err := NewTofuFS()
			//if err != nil {
			//log.Fatal(err)
			//}
			//if err = tfs.MountAndServe(c.Args().First(), c.Bool("readonly")); err != nil {
			//log.Fatal(err)
			//}
			return nil
		},
	}
}
