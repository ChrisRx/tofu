package client

import (
	//"bytes"
	"bufio"
	//"fmt"
	//"io"
	"log"
	"os"

	"golang.org/x/net/context"

	"github.com/ChrisRx/tofu/block"
	"github.com/ChrisRx/tofu/proto"
	"github.com/ChrisRx/tofu/volume"
)

type TofuClient struct {
	b *block.Client
	v *volume.Client
}

func NewTofuClient() *TofuClient {
	b := block.NewClient()
	v := volume.NewClient()
	return &TofuClient{
		b: b,
		v: v,
	}
}

func (t *TofuClient) GetFile(path string) ([]byte, error) {
	f, err := t.v.C.GetFile(context.Background(), &tofu.File{Path: path})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, block := range f.Blocks.Block {
		log.Println(block)
	}
	return nil, nil
}

func (t *TofuClient) PutFile(path string) {
	fs, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(fs)
	buf := make([]byte, 2048)
	f := &tofu.FileInfo{
		File: &tofu.File{Path: path},
		Blocks: &tofu.Blocks{
			Block: []*tofu.Block{},
		},
	}
	for {
		data, _ := reader.Read(buf)
		if data == 0 {
			break
		}
		block, err := t.b.PutBlock(buf)
		if err != nil {
			log.Fatal(err)
		}
		f.Blocks.Block = append(f.Blocks.Block, block)
	}
	a, err := t.v.C.PutFile(context.Background(), f)
	if err != nil {
		log.Println(err)
	}
	log.Println(a)
}
