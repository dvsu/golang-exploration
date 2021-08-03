package main

import "fmt"

func main() {
	// fix the following codes

	// toppings := []int{"black olives", "green peppers"}

	// var pizza [3]string
	// append(pizza, ...toppings)
	// pizza = append(toppings, "onions")
	// toppings = append(pizza, extra cheese)

	// fmt.Printf("pizza       : %s\n", pizza)

	toppings := []string{"black olives", "green peppers"}

	pizza := []string{}
	pizza = append(pizza, toppings...)
	pizza = append(pizza, "onions", "extra cheese")

	fmt.Printf("pizza: %s\n", pizza)
}
