package main

import (
	"fmt"

	"github.com/olezhek28/go-patterns/patterns/decorator/pizza"
)

func main() {
	olezhekPizza := &pizza.AnchovyTopping{
		Pizza: &pizza.TomatoTopping{
			Pizza: &pizza.PepperoniTopping{
				Pizza: &pizza.PepperoniTopping{
					Pizza: &pizza.ArtichokeTopping{
						Pizza: pizza.NewCheesePizza(),
					},
				},
			},
		},
	}

	fmt.Printf("Price of olezhek pizza: %v\n", olezhekPizza.GetPrice())
	fmt.Printf("Calories of olezhek pizza: %v\n", olezhekPizza.GetCalories())
}
