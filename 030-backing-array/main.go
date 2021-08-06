package main

import (
	"fmt"
	"sort"
)

func main() {
	// slice literal creates a hidden array and returns a slice that refers to that array
	// this hidden array is called backing array
	// backing array is stored separately from the slice
	names := []string{"Brad", "Caroline", "Amanda", "Daniel"}

	fmt.Printf("ptr:%p Original slice: %s\n", &names, names)

	// new slice whose elements are extracted from slicing existing slice,
	// will still point to the same backing array, even both have different slice header.
	// Different slice header = different pointer address
	names2 := names[0:2]

	fmt.Printf("ptr:%p New slice: %s \n", &names2, names2)

	// it also means, if we modify the original slice,
	// it will also modify the slice copy, because both
	// point to the same backing array

	// we will sort slice 'names'. Slice 'names2' will also be
	// sorted automatically

	sort.Strings(names)

	fmt.Printf("ptr:%p Original slice sorted: %s\n", &names, names)
	fmt.Printf("ptr:%p New slice (also sorted): %s\n", &names2, names2)

	// but there is also a way to create entirely separate slice from the original one
	// new slice 'newNames' will point to different memory address as it is entirely new slice
	var newNames []string
	newNames = append(newNames, names...)
	fmt.Printf("ptr:%p newNames slice: %s\n", &newNames, newNames)

	// we can also do it by using short declaration
	// first we create the nil slice, then append the original slice to the nil slice
	// Then, store all elements in new variable, 'otherNames'

	// otherNames := append([]string(nil), names...)

	// alternative syntax
	otherNames := append([]string{}, names...)

	fmt.Printf("ptr:%p otherNames slice: %s\n", &otherNames, otherNames)

	// now let's change an element in 'names' slice and prove the 'newNames' and 'otherNames' slices
	// are not modified

	names[3] = "Grant"

	fmt.Println("Change an element in 'names' slice")

	fmt.Printf("ptr:%p Recheck Original slice: %s\n", &names, names)
	fmt.Printf("ptr:%p Recheck newNames slice: %s\n", &newNames, newNames)
	fmt.Printf("ptr:%p Recheck otherNames slice: %s\n", &otherNames, otherNames)
}
