package main

import "fmt"

func main() {
	// We will try to embed a common type into
	// other types to mimic OOP composition concept.
	// We will refactor codes from part 56: type assertion

	// Because cryptocurrency and fiatcurrency have almost identical
	// struct fields and methods, a common type, 'currency' will be created
	// so both cryptocurrency and fiatcurrency can 'embed' the common type in
	// their struct.

	var (
		bitcoin  = cryptocurrency{currency{name: "Bitcoin", symbol: "BTC", rate: 47657.90}}
		ethereum = cryptocurrency{currency{name: "Ethereum", symbol: "ETH", rate: 3307.68}}
		litecoin = cryptocurrency{currency{name: "Litecoin", symbol: "LTC", rate: 182.79}}
		euro     = fiatcurrency{currency{name: "Euro", symbol: "EUR", rate: 1.18}}
		pound    = fiatcurrency{currency{name: "Pound sterling", symbol: "GBP", rate: 1.39}}
	)

	var currencies exchangerates

	// Note:
	// Go cannot automatically take a dynamic value's address
	// from an interface value because a non-pointer value is
	// non-addressable, a copy

	// the currency type has pointer receiver type, so we will pass the address
	// of each variable
	currencies = append(currencies, &bitcoin, &ethereum, &litecoin, &euro, &pound)
	fmt.Println("Initial rate")
	currencies.print()

	// the rate of all currency is up by 50%
	currencies.toTheMoon(0.5)

	fmt.Println("Final rate")
	currencies.print()

}