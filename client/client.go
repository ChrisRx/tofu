package client

import (
	"bytes"
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"tofu/proto"
)

const (
	address = "localhost:50051"
)

type TofuClient struct {
	conn *grpc.ClientConn
	c    tofu.BlockStoreClient
}

func NewTofuClient() *TofuClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := tofu.NewBlockStoreClient(conn)
	return &TofuClient{
		conn: conn,
		c:    c,
	}
}

func (t *TofuClient) GetBlock(hash string) ([]byte, error) {
	stream, err := t.c.GetBlock(context.Background(), &tofu.Block{Hash: hash})
	if err != nil {
		log.Fatalf("could not get block: %v", err)
	}
	bb := new(bytes.Buffer)
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		bb.Write(b.Data)
		log.Printf("%s", b.Data)
	}
	return bb.Bytes(), nil
}

// Break files into componets at top level THEN client/server
func (t *TofuClient) PutBlock() {
}

func (t *TofuClient) ListBlocks() []*tofu.Block {
	stream, err := t.c.ListBlocks(context.Background(), &tofu.ListBlocksRequest{})
	if err != nil {
		log.Fatalf("could not get block list: %v", err)
	}
	results := []*tofu.Block{}
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s", b.Hash)
		results = append(results, b)
	}
	return results
}

func (t *TofuClient) Close() {
	if t.conn != nil {
		t.conn.Close()
	}
}
