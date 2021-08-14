package main

import "fmt"

type fiatcurrency struct {
	name, symbol string
	rate         currency
}

func (f *fiatcurrency) print() {
	fmt.Printf("%s-USD: %9s\n", f.symbol, f.rate.string())
}