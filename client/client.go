package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

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
		data, err := ioutil.ReadFile(filepath.Join("data", block.Hash))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", data)
	}
	return nil, nil
}

func (t *TofuClient) PutFile(path string) {
	fs, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(fs)
	buf := make([]byte, 0, 2048)
	f := &tofu.FileInfo{
		File: &tofu.File{Path: path},
		Blocks: &tofu.Blocks{
			Block: []*tofu.Block{},
		},
	}
	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
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
