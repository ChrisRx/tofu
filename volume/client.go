package volume

import (
	"log"

	"google.golang.org/grpc"

	"github.com/ChrisRx/tofu/proto"
)

const (
	address = "localhost:50052"
)

type Client struct {
	conn *grpc.ClientConn
	C    tofu.VolumeStoreClient
}

func NewClient() *Client {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := tofu.NewVolumeStoreClient(conn)
	return &Client{
		conn: conn,
		C:    c,
	}
}
