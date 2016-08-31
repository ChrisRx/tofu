package block

import (
	"bytes"
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	"github.com/ChrisRx/tofu/proto"
)

const (
	address = "localhost:50051"
)

type Client struct {
	conn *grpc.ClientConn
	c    tofu.BlockStoreClient
}

func NewClient() *Client {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := tofu.NewBlockStoreClient(conn)
	return &Client{
		conn: conn,
		c:    c,
	}
}

func (t *Client) GetBlock(hash string) ([]byte, error) {
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
	}
	return bb.Bytes(), nil
}

func (t *Client) PutBlock(b []byte) (*tofu.Block, error) {
	stream, err := t.c.PutBlock(context.Background())
	if err := stream.Send(&tofu.BytesValue{Data: b}); err != nil {
		return nil, err
	}
	r, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (t *Client) ListBlocks() []*tofu.Block {
	stream, err := t.c.ListBlocks(context.Background(), &tofu.EmptyValue{})
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

func (t *Client) Close() {
	if t.conn != nil {
		t.conn.Close()
	}
}
