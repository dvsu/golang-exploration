package main

func main() {
	// Interface describes the expected behavior
	// The data type does not matter as long as
	// it has the expected behavior (method)
	// Interface is an abstract type
	// All non-interface types are concrete types
	// such as int, string, array, struct, slice,
	// map, func, etc.

	/*
				Interface syntax

			    interface name
				common convention
		        "-er" suffix
						|			 .------- interface type
						V            V
				type printer interface {
					print()
				}     ^
			          '-------------- describe the expected
			                          behavior (methods)

	*/

	// The bigger the interface the weaker the abstraction.
	// It is better to include as few methods as possible
	// because including many methods in the interface
	// method makes it very difficult to satisfy.

	var (
		bitcoin  = cryptocurrency{name: "Bitcoin", symbol: "BTC", rate: 47657.90}
		ethereum = cryptocurrency{name: "Ethereum", symbol: "ETH", rate: 3307.68}
		litecoin = cryptocurrency{name: "Litecoin", symbol: "LTC", rate: 182.79}
		euro     = fiatcurrency{name: "Euro", symbol: "EUR", rate: 1.18}
		pound    = fiatcurrency{name: "Pound sterling", symbol: "GBP", rate: 1.39}
	)

	var currencies exchangerates

	currencies = append(currencies, &bitcoin, &ethereum, &litecoin, &euro, &pound)
	currencies.print()

}