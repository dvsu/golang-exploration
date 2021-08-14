package main

import "fmt"

type cryptocurrency struct {
	name, symbol string
	rate         currency
}

func (c *cryptocurrency) print() {
	fmt.Printf("%s-USD: %9s\n", c.symbol, c.rate.string())
}