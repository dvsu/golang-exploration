package main

import (
	"fmt"
	"strings"
)

type country struct {
	name, capital string
}

type continent struct {
	area      int
	areaUnit  string
	countries []country
}

func main() {
	// struct values always get copied if
	// they are copied, passed to, or returned
	// from a function

	asia := continent{
		area:     44579000,
		areaUnit: "km2",
		countries: []country{
			{
				name:    "Japan",
				capital: "Tokyo",
			},
			{
				name:    "India",
				capital: "New Delhi"},
		},
	}

	fmt.Printf("%-12s%-12s\n", "Country", "Capital")
	fmt.Printf("%-24s\n", strings.Repeat("=", 24))
	for _, v := range asia.countries {

		v.name = "Changed"
		v.capital = "Changed"
		fmt.Printf("%-12s%-12s\n", v.name, v.capital)
	}

	fmt.Println("Recheck again after changing values to 'Changed'")
	fmt.Printf("%-12s%-12s\n", "Country", "Capital")
	fmt.Printf("%-24s\n", strings.Repeat("=", 24))
	for _, v := range asia.countries {
		fmt.Printf("%-12s%-12s\n", v.name, v.capital)
	}

	// Notice the real value is unchanged because the 'v' variable is actually
	// a copy of original one. Modifying the 'v' variable only changes
	// the value of the copied one.

	// what if we clone and change field values of both structs?
	newAsia := asia
	// Area of Asia is originally 44579000km2. 
	newAsia.area = 12345678
	asia.area = 42000000
	// Similarly, the old struct is not affected because 'newAsia'
	// is the clone of 'Asia'. Changing any field values of 'newAsia'
	// will only be applied to 'newAsia' and vice versa.
	fmt.Printf("Area of Asia: %d%s\n", asia.area, asia.areaUnit)
	fmt.Printf("Area of New Asia: %d%s\n", newAsia.area, newAsia.areaUnit)
}