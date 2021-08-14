package main

import "fmt"

func main() {
	// type assertion provides access to an interface value's
	// underlying concrete value

	/*
			syntax
		          .-------------- interface
		         /  .-------------------- data type
			     V  V
			t := i.(T)

			Type assertion can return 2 values, the underlying value and assertion result.
			true for successful assertion and false for failed assertion, i.e. different data type

			t, ok := i.(T)
	*/

	var i interface{} = "test"

	s, ok := i.(string)
	fmt.Println(s, ok) // true because i has underlying type of string

	f, ok := i.(float64)
	fmt.Println(f, ok) // false because i has underlying type of string

	var (
		bitcoin  = cryptocurrency{name: "Bitcoin", symbol: "BTC", rate: 47657.90}
		ethereum = cryptocurrency{name: "Ethereum", symbol: "ETH", rate: 3307.68}
		litecoin = cryptocurrency{name: "Litecoin", symbol: "LTC", rate: 182.79}
		euro     = fiatcurrency{name: "Euro", symbol: "EUR", rate: 1.18}
		pound    = fiatcurrency{name: "Pound sterling", symbol: "GBP", rate: 1.39}
	)

	var currencies exchangerates

	currencies = append(currencies, &bitcoin, &ethereum, &litecoin, &euro, &pound)
	fmt.Println("Initial rate")
	currencies.print()


	// What if there is a method that is only available in specific type
	// and we want to call it?

	// The answer is, we can use type assertion on the interface 
	// to check whether a variable that we are interested in 
	// has the right data type

	// Currently, the interface contains data with 2 different data types, 
	// cryptocurrency and fiatcurrency. cryptocurrency has another method
	// named toTheMoon(), which updates the rate of a cryptocurrency 
	// by a specified factor. If we try to loop over the elements and
	// call the method on each element, it will cause an error because
	// fiatcurrency type does not have toTheMoon() method. This is when 
	// we can use type assertion to check whether the data type of an
	// element matches with the one we are interested in. If it matches,
	// we can safely call the method.


	// the rate of all cryptocurrency is up by 50%
	currencies.toTheMoon(0.5)

	fmt.Println("Final rate")
	currencies.print()
}