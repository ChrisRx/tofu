package store

import (
	"errors"
	"time"

	"github.com/ChrisRx/tofu/proto"
)

var (
	ErrNodeNotDirectory = errors.New("Node is not directory")
)

type Node struct {
	Path  string
	Value string

	Parent   *Node
	Children map[string]*Node

	created  int64
	modified int64
}

func newKV(path string, value string) (n *Node) {
	cur := time.Now().Unix()
	n = Node{
		Path:     path,
		Value:    value,
		created:  cur,
		modified: cur,
	}
	return
}

func newDir(path string, value string) (n *Node) {
	n = newKV(path, "")
	n.Children = make(map[string]*Node)
	return
}

func (n *Node) IsDir() bool {
	return n.Children != nil
}

func (n *Node) Created() int64 {
	return n.created
}

func (n *Node) Modified() int64 {
	return n.modified
}

func (n *Node) Get(v string) (string, error) {
	return n.Value, nil
}

func (n *Node) Set(v string) error {
	n.Value = v
	n.modified = time.Now().Unix()
	return nil
}

func (n *Node) List() ([]*Node, error) {
	if !n.isDir {
		return nil, ErrNodeNotDirectory
	}
	nodes := make([]*Node, len(n.Children))
	for i, node := range n.Children {
		nodes[i] = node
	}
	return nodes, nil
}
