package main

func main() {
	// Method can also be attached to any types, such as basic types (int, string),
	// composite types (array, struct, slice, map, chan), or even function
	// For slice, map, chan, and func, do not use pointer receiver because
	// they have their own pointer.
	// So, it does not have to be a struct type.

	var (
		bitcoin  = cryptocurrency{name: "Bitcoin", symbol: "BTC", rate: 47657.90}
		ethereum = cryptocurrency{name: "Ethereum", symbol: "ETH", rate: 3307.68}
		litecoin = cryptocurrency{name: "Litecoin", symbol: "LTC", rate: 182.79}
	)

	var cryptocurrencies []*cryptocurrency
	cryptocurrencies = append(cryptocurrencies, &bitcoin, &ethereum, &litecoin)

	rates := exchangerates(cryptocurrencies)

	rates.print()
}
