package main

import "fmt"

type exchangerate interface {
	print()
	toTheMoon(float64)
}

type exchangerates []exchangerate

func (ex exchangerates) print() {
	if len(ex) == 0 {
		fmt.Println("There is no currencies available")
		return
	}

	for _, c := range ex {
		c.print()
	}
}

func (ex exchangerates) toTheMoon(f float64) {

	for _, c := range ex {
		c.toTheMoon(f)
	}
} 