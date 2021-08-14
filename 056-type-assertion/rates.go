package main

import "fmt"

type printer interface {
	print()
}

type exchangerates []printer

func (ex exchangerates) print() {
	if len(ex) == 0 {
		fmt.Println("There is no currencies available")
		return
	}

	for _, t := range ex {
		t.print()
	}
}

func (ex exchangerates) toTheMoon(f float64) {

	// We do not really have to declare an interface at
	// package level. It can also be declared inside a function
	type toTheMooner interface {
		toTheMoon(float64)
	}

	for _, c := range ex {
		// type assertion
		// t, isCrypto := c.(*cryptocurrency)
		// if !isCrypto {
		// 	continue
		// }
		
		// t.toTheMoon(f)

		// alternatively, to check if the element
		// has method named toTheMoon(f)
		if t, isCrypto := c.(toTheMooner); isCrypto {
			t.toTheMoon(f)
		}
	}
} 