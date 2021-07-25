package main

import "fmt"

func main() {
	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet, isTrue, temp = "Mars", true, 19.5

	fmt.Printf("Air is good on %v\nIt's %v\nIt is %v degrees\n", planet, isTrue, temp)
}
