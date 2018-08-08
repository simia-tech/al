// +build js

package local

import (
	"fmt"
	"syscall/js"

	"github.com/simia-tech/al/storage"
)

type Storage struct {
	storage.Interface
	Name      string
	lsElement js.Value
}

func NewStorage(name string) (*Storage, error) {
	lsElement := js.Global().Get("localStorage")
	if lsElement.Type() != js.TypeObject {
		return nil, fmt.Errorf("could not find localStorage object (found %s)", lsElement.Type())
	}
	return &Storage{
		Name:      name,
		lsElement: lsElement,
	}, nil
}

func (s *Storage) Get(key string) (string, error) {
	item := s.lsElement.Call("getItem", s.key(key))
	if item.Type() == js.TypeString {
		return item.String(), nil
	}
	return "", nil
}

func (s *Storage) Set(key, value string) error {
	s.lsElement.Call("setItem", s.key(key), value)
	return nil
}

func (s *Storage) key(key string) string {
	if s.Name == "" {
		return key
	}
	return s.Name + "_" + key
}
