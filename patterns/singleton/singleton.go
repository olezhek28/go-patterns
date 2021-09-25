package singleton

import (
	"fmt"
	"sync"
)

var (
	once           sync.Once
	singleInstance *single
)

type single struct {
	name string
}

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Create...")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Already")
	}

	return singleInstance
}

func (s *single) SetName(name string) {
	s.name = name
}

func (s *single) GetName() string {
	return s.name
}