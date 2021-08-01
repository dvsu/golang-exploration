package main

import (
	"fmt"
	"strings"
)

func main() {
	// refactor exercise 071 using array literal
	names := [...]string{
		"America",
		"Asia",
		"Europe",
	}

	distances := [...]int{
		10,
		29,
		38,
		47,
		56,
	}

	data := [...]byte{
		'P',
		'L',
		'U',
		'T',
		'O',
	}

	ratios := [...]float64{
		2.7182818284,
	}

	alives := [...]bool{
		true,
		false,
		false,
		true,
	}

	// when the element of array is empty, don't do this
	// zero := [...]byte{}
	// instead
	var zero [0]byte

	// let's print the elements inside each array
	t1 := "'names' array"
	fmt.Printf("%s %s %[1]s\n", strings.Repeat("=", (50-len(t1))/2), t1)
	for i, e := range names {
		fmt.Printf("names[%d] %q\n", i, e)
	}

	t2 := "'distances' array"
	fmt.Printf("%s %s %[1]s\n", strings.Repeat("=", (50-len(t2))/2), t2)
	for i, e := range distances {
		fmt.Printf("distances[%d] %d\n", i, e)
	}

	t3 := "'data' array"
	fmt.Printf("%s %s %[1]s\n", strings.Repeat("=", (50-len(t3))/2), t3)
	for i, e := range data {
		fmt.Printf("data[%d] %d\n", i, e)
	}

	t4 := "'ratios' array"
	fmt.Printf("%s %s %[1]s\n", strings.Repeat("=", (50-len(t4))/2), t4)
	for i, e := range ratios {
		fmt.Printf("ratios[%d] %f\n", i, e)
	}

	t5 := "'alives' array"
	fmt.Printf("%s %s %[1]s\n", strings.Repeat("=", (50-len(t5))/2), t5)
	for i, e := range alives {
		fmt.Printf("alives[%d] %t\n", i, e)
	}

	t6 := "'zero' array"
	fmt.Printf("%s %s %[1]s\n", strings.Repeat("=", (50-len(t6))/2), t6)
	for i, e := range zero {
		fmt.Printf("zero[%d] %q\n", i, e)
	}
}
