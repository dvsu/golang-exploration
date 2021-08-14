package main

import "fmt"

type cryptocurrency struct {
	name string
	symbol string
	rate currency
}

func (cc *cryptocurrency) print() {
	fmt.Printf("%s-USD %9s\n", cc.symbol, cc.rate.string())
}
