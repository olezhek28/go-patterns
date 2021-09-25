package singleton

import (
	"fmt"
	"sync"
)

var (
	once     sync.Once
	instance *single = nil
)

type Singleton interface {
	SetName(name string)
	GetName() string
}

type single struct {
	name string
}

func GetInstance() Singleton {
	once.Do(
		func() {
			fmt.Println("Create...")
			instance = new(single)
		})

	return instance
}

func (s *single) SetName(name string) {
	s.name = name
}

func (s *single) GetName() string {
	return s.name
}
