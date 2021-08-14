package main

import "fmt"

type currency struct {
	name, symbol string
	rate         dollar
}

func (c *currency) print() {
	fmt.Printf("%s-USD: %9s\n", c.symbol, c.rate.string())
}

func (c *currency) toTheMoon(f float64) {	
	c.rate *= dollar(1 + f)
}
