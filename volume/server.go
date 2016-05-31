package volume

import (
	//"bytes"
	"errors"
	//"fmt"
	//"io"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/ChrisRx/tofu/block"
	"github.com/ChrisRx/tofu/proto"
)

var (
	ErrFileNotFound = errors.New("File was not found")
)

type Volume struct {
	b *block.Client

	files map[string]*tofu.FileInfo
}

func New() *Volume {
	b := block.NewClient()
	f := make(map[string]*tofu.FileInfo)
	return &Volume{
		b:     b,
		files: f,
	}
}

func (v *Volume) GetFile(ctx context.Context, file *tofu.File) (*tofu.FileInfo, error) {
	log.Println(v.files)
	if val, ok := v.files[file.Path]; ok {
		return val, nil
	}
	return nil, ErrFileNotFound
}

func (v *Volume) PutFile(ctx context.Context, file *tofu.FileInfo) (*tofu.FileInfo, error) {
	log.Println("Hur")
	log.Println(file)
	v.files[file.File.Path] = file
	return file, nil
}

func (v *Volume) ListFiles(emptyValue *tofu.EmptyValue, stream tofu.VolumeStore_ListFilesServer) error {
	return nil
}

func (v *Volume) Run() {
	var opts []grpc.ServerOption
	port := ":50052"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(opts...)
	tofu.RegisterVolumeStoreServer(s, v)
	s.Serve(lis)
}
