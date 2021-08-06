package main

import (
	"fmt"
)

func main() {
	// capacity is the length of backing array
	// it is impossible to extend a slice beyond its capacity

	// length vs capacity
	// length = length of a slice
	// capacity = length of backing array from the beginning of the first element of slice

	sources := []int{10, 29, 38, 47}
	chopped := sources[1:3]
	fmt.Printf("'sources' slice. Length: %d Capacity: %d Values: %d\n", len(sources), cap(sources), sources)
	// length: 2 (because it contains 29, 38) capacity: 3 (because it can only see 29, 38, 47)
	fmt.Printf("'chopped' slice. Length: %d Capacity: %d Values: %d\n", len(chopped), cap(chopped), chopped)
	rechopped := sources[:2]
	// length: 2 (because it contains 10, 29) capacity: 4 (because it can see the entire backing array 10, 29, 38, 47)
	// but can only access 10 and 29
	fmt.Printf("'rechopped' slice. Length: %d Capacity: %d Values: %d\n", len(rechopped), cap(rechopped), rechopped)

	words := []string{"the", "world", "out", "there", "is", "a", "good", "place"}
	fmt.Printf("'words' slice. Length: %d Capacity: %d ptr (backing array): %p Values: %s\n", len(words), cap(words), words, words)
	words2 := words[:0]
	// because it is pointing to the beginning of backing array,
	// the length is reduced to zero, but the capacity remains the same, 5.
	fmt.Printf("'words2' slice. Length: %d Capacity: %d ptr (backing array): %p Values: %s\n", len(words2), cap(words2), words2, words2)

	words3 := words[0:]
	// because it is pointing to the beginning of backing array,
	// the length is reduced to zero, but the capacity remains the same, 5.
	fmt.Printf("'words3' slice. Length: %d Capacity: %d ptr (backing array): %p Values: %s\n", len(words3), cap(words3), words3, words3)

	// nil slice
	// pointer value: zero, length: zero, capacity: zero
	var weeks []int
	fmt.Printf("Pointer value of 'weeks' backing array: %p\n", weeks)
	fmt.Printf("Address of 'weeks' slice: %p\n", &weeks)

	// empty slice
	// pointer value: non-zero, length: zero, capacity: zero
	// pointer value is not zero because Go assigns dummy backing array to empty slices
	months := []string{}
	fmt.Printf("Pointer value of 'months' backing array: %p\n", months)
	fmt.Printf("Address of 'months' slice: %p\n", &months)

	// if we create another empty slice, the pointer value will have the same value
	// because empty slices share the same backing array, regardless the element type
	years := []int{}
	fmt.Printf("Pointer value of 'years' backing array: %p\n", years)
	fmt.Printf("Address of 'years' slice: %p\n", &years)

	// once the slice is filled with some elements, Go will create the real backing array
	// for the slice. It also means it will have different pointer value as compared to
	// pointer value of empty backing array
	// However, the address of slice is still the same
	years = []int{1975, 1980, 1999, 2010}
	fmt.Printf("Real pointer value of 'years' backing array: %p\n", years)
	fmt.Printf("Address of 'years' slice: %p\n", &years)
}
