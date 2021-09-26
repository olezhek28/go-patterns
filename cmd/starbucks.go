package main

import (
	"fmt"

	"github.com/olezhek28/go-patterns/patterns/decorator/starbucks"
)

func main() {
	var coffee starbucks.Beverage = starbucks.NewEspresso()
	coffee = &starbucks.Milk{Beverage: coffee}
	coffee = &starbucks.Rum{Beverage: coffee}
	coffee = &starbucks.Chocolate{Beverage: coffee}

	fmt.Printf("Your order: %v\n", coffee.GetDescription())
	fmt.Printf("From you: %v $\n", coffee.Cost())
}
