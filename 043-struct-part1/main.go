package main

import "fmt"

func main() {
	// struct is similar to class in OOP languages.
	// struct is fixed at compile-time, so
	// the fields and its type cannot be changed after compilation.
	// Field values, however, can be changed during runtime.
	// Field name has to be unique, similar to 'key name' in map.
	// Every uninitialized struct field will get a zero-value
	// depending on its type.

	type country struct {
		capital, continent string
		population         int
	}

	var (
		Germany country
		Canada  country
	)

	fmt.Println("Empty structs")
	fmt.Printf("Germany: %#v\n", Germany)
	fmt.Printf("Canada: %#v\n", Canada)

	Germany = country{
		capital: "Berlin",
		continent: "Europe",
	}

	Japan := country{
		capital: "Tokyo",
		continent: "Asia",
		population: 125360000,
	}

	Canada.population = 38131104
	fmt.Println("Structs + values")
	fmt.Printf("Germany: %#v\n", Germany)
	fmt.Printf("Canada: %#v\n", Canada)
	fmt.Printf("Japan: %#v\n", Japan)

	fmt.Println("Accessing struct field")
	fmt.Printf("%s is the capital of Germany.\n", Germany.capital)	
}