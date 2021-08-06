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
	names2 := names[1:3]

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

	fmt.Printf("ptr (header):%p ptr (backing array): %p otherNames slice: %s\n", &otherNames, otherNames, otherNames)

	// now let's change an element in 'names' slice and prove the 'newNames' and 'otherNames' slices
	// are not modified

	names[3] = "Grant"

	fmt.Println("Change an element in 'names' slice")

	fmt.Printf("ptr (header):%p ptr (backing array):%p Recheck Original slice (names): %s\n", &names, names, names)
	fmt.Println(`Backing array address of 'names2' slice is 2 bytes (16 bits) apart from 'names'. However, they still share the same backing array 
because 'names2' [1:3] is a subset of 'name' [:], i.e. starts from the second element of 'name'`)
	fmt.Printf("ptr (header):%p ptr (backing array):%p Recheck names2 slice: %s\n", &names2, names2, names2)
	fmt.Printf("ptr (header):%p ptr (backing array):%p Recheck newNames slice: %s\n", &newNames, newNames, newNames)
	fmt.Printf("ptr (header):%p ptr (backing array):%p Recheck otherNames slice: %s\n", &otherNames, otherNames, otherNames)

	fmt.Println("Declare a new slice 'copiedNames' whose values are copied from 'names' slice")
	copiedNames := names
	fmt.Println("Reduce the capacity of 'copiedNames' slice to one")
	fmt.Printf("ptr (header):%p ptr (backing array):%p Recheck copiedNames slice: %s\n", &copiedNames, copiedNames, copiedNames)
	copiedNames = copiedNames[1:cap(copiedNames)]
	fmt.Printf("ptr (header):%p ptr (backing array):%p Recheck copiedNames slice: %s\n", &copiedNames, copiedNames, copiedNames)
	copiedNames = copiedNames[1:cap(copiedNames)]
	fmt.Printf("ptr (header):%p ptr (backing array):%p Recheck copiedNames slice: %s\n", &copiedNames, copiedNames, copiedNames)
	copiedNames = copiedNames[1:cap(copiedNames)]
	fmt.Printf("ptr (header):%p ptr (backing array):%p Recheck copiedNames slice: %s\n", &copiedNames, copiedNames, copiedNames)
	fmt.Println("Now check on the original 'names' slice")
	fmt.Printf("ptr (header):%p ptr (backing array):%p Recheck Original slice (names): %s\n", &names, names, names)

	names2 = append(names2, "Georgia")
	fmt.Println("Append 'Georgia' as third element to 'names2' slice")
	fmt.Printf("ptr (header):%p ptr (backing array):%p Recheck Original slice (names2): %s\n", &names2, names2, names2)
	fmt.Println("Last element of 'names' slice is also changed to 'Georgia' because both share the same backing array")
	fmt.Printf("ptr (header):%p ptr (backing array):%p Recheck Original slice (names): %s\n", &names, names, names)

}
