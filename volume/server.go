package volume

import (
	"errors"
	"log"
	"net"
	"time"

	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/ChrisRx/tofu/proto"
)

var (
	ErrFileNotFound      = errors.New("File was not found.")
	ErrConnectionRefused = errors.New("Unable to connect to etcd.")
	ErrOperationFailed   = errors.New("Operation failed.")

	DefaultTimeout = 5 * time.Second
)

type Volume struct {
	*clientv3.Client
}

func New() *Volume {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	return &Volume{
		Client: cli,
	}
}

func (v *Volume) GetFile(ctx context.Context, file *tofu.File) (*tofu.FileInfo, error) {
	log.Printf("GetFile: %s\n", file.Path)
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	resp, err := v.Get(ctx, file.Path)
	cancel()
	if err != nil {
		return nil, err
	}
	for _, ev := range resp.Kvs {
		f := &tofu.FileInfo{}
		err = f.Unmarshal(ev.Value)
		if err != nil {
			return nil, err
		}
		return f, nil
	}
	return nil, ErrFileNotFound
}

func (v *Volume) PutFile(ctx context.Context, file *tofu.FileInfo) (*tofu.FileInfo, error) {
	data, err := file.Marshal()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	_, err = v.Put(ctx, file.File.Path, string(data))
	cancel()
	if err != nil {
		return nil, err
	}
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
