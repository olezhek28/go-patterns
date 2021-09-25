package main

import (
	"fmt"

	"github.com/olezhek28/go-patterns/patterns/singleton"
)

func main() {
	obj1 := singleton.GetInstance()
	obj1.SetName("first")
	fmt.Println(obj1.GetName())

	obj2 := singleton.GetInstance()
	obj2.SetName("second")
	fmt.Println(obj1.GetName())
	fmt.Println(obj2.GetName())
}

