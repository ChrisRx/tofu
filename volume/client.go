package volume

import (
	//"bytes"
	//"fmt"
	//"io"
	"log"

	//"golang.org/x/net/context"
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

//func (t *TofuClient) GetFile(path string) ([]byte, error) {
//}
//func (t *TofuClient) PutFile(path string) {
//}
