package block

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/satori/go.uuid"
	"google.golang.org/grpc"

	"github.com/ChrisRx/tofu/proto"
)

type BlockStore struct {
	root string
}

func NewBlockStore() *BlockStore {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		log.Printf("Data dir does not exist, creating ...\n")
		os.Mkdir("data", 0777)
	}
	return &BlockStore{
		root: "data",
	}
}

func (b *BlockStore) GetBlock(block *tofu.Block, stream tofu.BlockStore_GetBlockServer) error {
	log.Printf("Got block %s\n", block.Hash)
	p := filepath.Join(b.root, block.Hash)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		log.Printf("Block %s does not exists.", block.Hash)
	}
	data, err := ioutil.ReadFile(p)
	if err != nil {
		log.Println(err)
	}
	if err := stream.Send(&tofu.BytesValue{Data: data}); err != nil {
		return err
	}
	return nil
}

func (b *BlockStore) PutBlock(stream tofu.BlockStore_PutBlockServer) error {
	bb := new(bytes.Buffer)
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			u := uuid.NewV4()
			result := &tofu.Block{Hash: u.String()}
			p := filepath.Join(b.root, result.Hash)
			if _, err := os.Stat(p); err == nil {
				log.Printf("Block %s already exists.", result.Hash)
			}
			err := ioutil.WriteFile(p, bb.Bytes(), 0644)
			if err != nil {
				log.Println(err)
			}
			log.Printf("PutBlock: %s\n", result.Hash)
			return stream.SendAndClose(result)
		}
		if err != nil {
			log.Fatal(err)
		}
		n, err := bb.Write(r.Data)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Added %d bytes\n", n)
	}
	return nil
}

func (b *BlockStore) ListBlocks(emptyValue *tofu.EmptyValue, stream tofu.BlockStore_ListBlocksServer) error {
	files, err := ioutil.ReadDir(b.root)
	if err != nil {
		log.Println("Hi: ", err)
	}
	for _, f := range files {
		if err := stream.Send(&tofu.Block{Hash: f.Name()}); err != nil {
			return err
		}
	}
	return nil
}

func (b *BlockStore) Run() {
	var opts []grpc.ServerOption
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(opts...)
	tofu.RegisterBlockStoreServer(s, b)
	s.Serve(lis)
}
