package main

import (
	"fmt"
	"os"

	"golang.org/x/net/context"

	"bazil.org/fuse"
	fusefs "bazil.org/fuse/fs"
)

type Dir struct {
	Name    string
	Files   []File
	Folders map[string]*Dir

	fs *TofuFS
}

func NewDir(rootDir *Dir, fs *TofuFS) *Dir {
	return &Dir{
		Name:    rootDir.Name,
		Files:   rootDir.Files,
		Folders: rootDir.Folders,
		fs:      fs,
	}
}

func (d *Dir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = 1
	a.Uid = uint32(os.Getuid())
	a.Gid = uint32(os.Getgid())
	a.Mode = os.ModeDir | 0755
	return nil
}

func (d *Dir) Lookup(ctx context.Context, req *fuse.LookupRequest, resp *fuse.LookupResponse) (fusefs.Node, error) {
	fmt.Printf("DirLookup: FolderName(%s) RequestName(%s)\n", d.Name, req.Name)
	return nil, fuse.ENOENT
}

func (d *Dir) Create(ctx context.Context, req *fuse.CreateRequest, resp *fuse.CreateResponse) (fusefs.Node, fs.Handle, error) {
}

func (d *Dir) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (fusefs.Handle, error) {
	fmt.Printf("DirOpen: Name(%s) RDir(%t) \n", d.Name, req.Dir)
	if d.Name != d.fs.cur {
		fmt.Printf("Changing directories: %s -> %s\n", d.fs.cur, d.Name)
		d.fs.cur = d.Name
		//return newDir, nil
	}
	return d, nil
}

func (d *Dir) Mkdir(ctx context.Context, req *fuse.MkdirRequest) (fusefs.Node, error) {
	if _, ok := d.Folders[req.Name]; ok {
		fmt.Printf("Directory %s already exists\n", req.Name)
		return nil, fuse.ENOENT
	}
	d.Folders[req.Name] = NewDir(&Dir{Name: req.Name}, d.fs)
	return d.Folders[req.Name], nil
}

func (d *Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	fmt.Printf("DirRead\n")
	l := []fuse.Dirent{}
	dirs := []fuse.Dirent{}
	for _, v := range d.Folders {
		dir := fuse.Dirent{
			Name: v.Name,
			Type: fuse.DT_Dir,
		}
		dirs = append(dirs, dir)
	}
	files := make([]fuse.Dirent, len(d.Files))
	for i, s := range d.Files {
		files[i].Name = s.Name
		files[i].Type = fuse.DT_Dir
	}

	l = append(l, dirs...)
	l = append(l, files...)

	return l, nil
}
