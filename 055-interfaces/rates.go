package main

import "fmt"

type printer interface {
	print()
}

type exchangerates []printer

func (cc exchangerates) print() {
	if len(cc) == 0 {
		fmt.Println("Cryptocurrency list is not available")
		return
	}

	for _, c := range cc {
		c.print()
	}
}
