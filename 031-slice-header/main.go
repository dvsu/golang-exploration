package main

import (
	"fmt"
	"strings"
	"unsafe"
)

type (
	collection  [5]string
	scollection []string
)

func main() {
	// slice header is a tiny data structure that describe its backing array.
	// It contains pointer, length, and capacity information
	// It is used by Go runtime, but not visible to user

	// nil slice does not have a backing array, but it has a slice header
	// with each value (pointer, length, capacity) equals to zero

	// passing an array to a function will create
	// a copy of the array in separate memory location
	// it will not modify the original one
	input := collection{"one", "two", "three", "ninety", "seventy"}
	fmt.Printf("ptr: %p Original array: %s\n", &input, input)
	processArray(input)
	processArray2(input)
	fmt.Printf("ptr: %p Recheck original array: %s\n", &input, input)
	fmt.Printf("%s\n", strings.Repeat("=", 80))

	// however, if an element of an array is modified locally
	// no copying occurs, so backing array and address remain the same
	input2 := collection{"four", "five", "six", "fifty", "thirty"}
	fmt.Printf("ptr: %p Original array2: %s\n", &input2, input2)
	input2[2] = "zero"
	fmt.Printf("ptr: %p Array2 processed locally: %s\n", &input2, input2)
	fmt.Printf("ptr: %p Recheck original array2: %s\n", &input2, input2)
	fmt.Printf("%s\n", strings.Repeat("=", 80))

	// Passing a slice to a function will create a new slice header
	// with new pointer address (notice the change in pointer address).
	// However, both still point to the same backing array.
	// As the result, modifying a slice element inside
	// a function still also modifies the original slice value.
	input3 := scollection{"seven", "eight", "nine", "twenty"}
	fmt.Printf("ptr: %p Original slice: %s\n", &input3, input3)
	processSlice(input3)
	fmt.Printf("ptr: %p Recheck original slice: %s\n", &input3, input3)

	// let's compare the size of array and slice after modification inside function
	// note: do not use unsafe.Sizeof() in production code
	fmt.Printf("Size of array after modification: %d\n", unsafe.Sizeof(input))
	fmt.Printf("Size of slice after modification: %d\n", unsafe.Sizeof(input3))

	var backingArraySize int

	for _, e := range input3 {
		backingArraySize += int(unsafe.Sizeof(e))
	}

	fmt.Printf("Size of slice's backing array after modification: %d bytes\n", backingArraySize)

	// size of array after modification = 80 bytes (16bytes * 5 array elements)
	// size of slice after modification = 24 bytes (8bytes * 3 header values) // regardless the quantity of elements
	// because the slice elements are stored separately in backing array
	// size of slice's backing array = 64 bytes (16 bytes * 4 backing array elements)

}

func processArray(input collection) {

	input[2] = "five"
	fmt.Printf("ptr: %p Array processed in function: %s\n", &input, input)
}

func processArray2(input collection) {

	input[3] = "secret"
	fmt.Printf("ptr: %p Array processed in second function: %s\n", &input, input)
}

func processSlice(input scollection) {

	input[2] = "ten"
	fmt.Printf("ptr: %p Slice processed in function: %s\n", &input, input)
}
