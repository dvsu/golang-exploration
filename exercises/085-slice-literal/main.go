package main

import "fmt"

func main() {
	// continuation of exercise 084
	// assign some values to these slices

	names := []string{"Abraham", "Bernard", "Charlie"}
	distances := []int{10, 89, 56, 47, 38}
	data := []uint8{'C', 'O', 'D', 'E'}
	ratios := []float64{2.7182818284}
	alives := []bool{true, false, true, true, false}

	fmt.Printf("names: %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data: %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios: %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives: %T %d %t\n", alives, len(alives), alives == nil)
}
