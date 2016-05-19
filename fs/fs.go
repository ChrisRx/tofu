package main

import (
	"log"

	"bazil.org/fuse"
	fusefs "bazil.org/fuse/fs"
)

var (
	programName = "TofuFS"
)

type TofuFS struct {
	mountDir string
	root     *Dir
	cur      string
}

func NewTofuFS() (*TofuFS, error) {
	r := &TofuFS{}
	return r, nil
}

func (fs *TofuFS) Root() (fusefs.Node, error) {
	if fs.root == nil {
		fs.root = NewDir(&example, fs)
	}
	return fs.root, nil
}

func (fs *TofuFS) MountAndServe(mountpoint string, readonly bool) error {
	fs.mountDir = mountpoint
	mountOpts := []fuse.MountOption{
		fuse.FSName(programName),
		fuse.Subtype(programName),
		fuse.VolumeName(programName),
		fuse.LocalVolume(),
	}
	if readonly {
		mountOpts = append(mountOpts, fuse.ReadOnly())
	}
	conn, err := fuse.Mount(mountpoint, mountOpts...)
	if err != nil {
		return err
	}
	defer conn.Close()

	OnInterrupt(func() {
		err := fuse.Unmount(mountpoint)
		if err != nil {
			log.Fatal(err)
		}
		conn.Close()
	})

	if err = fusefs.Serve(conn, fs); err != nil {
		return err
	}

	<-conn.Ready
	if err = conn.MountError; err != nil {
		return err
	}

	return nil
}
