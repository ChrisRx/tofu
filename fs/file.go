package main

import (
	"os"

	"bazil.org/fuse"
	"golang.org/x/net/context"
)

type File struct {
	Name string

	content []byte
}

func NewFile(name string, content []byte) *File {
	if content == nil {
		content = []byte("Dis file tho...\n")
	}
	return &File{
		Name:    name,
		content: content,
	}
}

func (f *File) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = 2
	a.Mode = 0444
	a.Uid = uint32(os.Getuid())
	a.Gid = uint32(os.Getgid())
	a.Size = uint64(len(f.content))
	return nil
}

func (f *File) ReadAll(ctx context.Context) ([]byte, error) {
	return f.content, nil
}
