package store

import (
	"sync"

	"github.com/ChrisRx/tofu/proto"
)

type Store struct {
	Root *Node

	mu sync.RWMutex
}

func New() *Store {
	return &Store{
		Root: newDir("/"),
	}
}

func (s *Store) Get() {
}
func (s *Store) Set() {
}

func (s *Store) Create() {
}
func (s *Store) Delete() {
}

func (s *Store) Save() {
}
