package main

import "fmt"

type customtype int

var x customtype
var y int

func main() {

	fmt.Printf("Value of variable x=%d Data type '%T'\n", x, x)
	x = 20
	fmt.Printf("Assigning x with a new value %d\n", x)
	y = int(x)
	fmt.Printf("Value of variable y=%d with data type '%T'\n", y, y)
}
