package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		names     [3]string
		distances [5]int
		data      [5]byte
		ratios    [1]float64
		alives    [4]bool
		zero      [0]byte
	)

	names[0] = "America"
	names[1] = "Asia"
	names[2] = "Europe"

	distances[0] = 10
	distances[1] = 29
	distances[2] = 38
	distances[3] = 47
	distances[4] = 56

	data[0] = 'P'
	data[1] = 'L'
	data[2] = 'U'
	data[3] = 'T'
	data[4] = 'O'

	ratios[0] = 2.7182818284

	alives[0] = true
	alives[1] = false
	alives[2] = false
	alives[3] = true

	// make sure the linter does not complain that zero is unused
	_ = zero

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
