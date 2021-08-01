package main

import "fmt"

func main() {
	// continuation of exercise 084
	// assign some values to these slices

	// arrays
	names := [3]string{"Abraham", "Bernard", "Charlie"}
	distances := [5]int{10, 89, 56, 47, 38}
	data := [4]uint8{'C', 'O', 'D', 'E'}
	ratios := [1]float64{2.7182818284}
	alives := [5]bool{true, false, true, true, false}
	zero := [0]byte{}

	// to slices
	fmt.Printf("names: %[1]T %[1]q\n", names[:])
	fmt.Printf("distances: %[1]T %[1]d\n", distances[:])
	fmt.Printf("data: %[1]T %[1]d\n", data[:])
	fmt.Printf("ratios: %[1]T %.3[1]f\n", ratios[:])
	fmt.Printf("alives: %[1]T %[1]t\n", alives[:])
	fmt.Printf("zero: %[1]T %[1]d\n", zero[:])
}
