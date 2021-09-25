package singleton

import (
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
	sync.RWMutex
}

func GetInstance() Singleton {
	once.Do(
		func() {
			instance = new(single)
		})

	return instance
}

func (s *single) SetName(name string) {
	s.Lock()
	defer s.Unlock()

	s.name = name
}

func (s *single) GetName() string {
	s.RLock()
	defer s.RUnlock()

	return s.name
}
