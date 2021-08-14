package main

import "fmt"

type computer struct {
	brand *string
}

func main() {
	// expected output
	// brand: apple

	// before
	// var c *computer
	// change(c, "apple")
	// fmt.Printf("brand: %s\n", c.brand)

	// assign variable c with an empty value instead of nil
	c := &computer{}
	change(c, "apple")

	// then extract the string value using *c.brand
	// which is equivalent to the value stored in
	// memory address 'brand'
	fmt.Printf("brand: %s\n", *c.brand)
}

func change(c *computer, brand string) {
	// before
	// (*c.brand) = brand

	// store the brand address to c.brand which stores
	// pointer value. It lets the 'brand' know where
	// to look for the actual value
	c.brand = &brand
}