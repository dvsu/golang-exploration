package main

import "fmt"

type cryptocurrency struct {
	name, symbol string
	rate         currency
}

func (cc *cryptocurrency) print() {
	fmt.Printf("%s-USD: %9s\n", cc.symbol, cc.rate.string())
}

func (cc *cryptocurrency) toTheMoon(f float64) {	
	cc.rate *= currency(1 + f)
}
