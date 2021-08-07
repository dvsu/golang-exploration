package main

import (
	"fmt"
	"strings"
)

func main() {
	// make can be used to initializes and returns a slice with given length and capacity

	// Syntax:
	// make({type of slice}, {length and capacity})
	// make({type of slice}, {length}, {capacity})

	// 'make' will preallocate backing array with specified length and capacity
	// so, it is useful to prevent re-creation of backing array.
	// Instead of creating a new backing array every time the size grows,
	// we can predefine the size at the beginning.

	// if we know the length of the slice that we need, it is recommended to use 'make'

	devices := []string{"mouse", "keyboard", "monitor", "speaker", "microphone"}

	fmt.Println("make()")
	// let's copy 'devices' slice and modify each element to Uppercase
	// we'll create a new slice with type []string, length of 0 and capacity of 5
	copied := make([]string, 0, len(devices))
	fmt.Printf("Length: %d Capacity: %d %p %s\n", len(copied), cap(copied), copied, copied)

	// then modify the elements of 'devices' append the elements one by one
	for _, device := range devices {
		copied = append(copied, strings.ToUpper(device))
		fmt.Printf("Length: %d Capacity: %d %p %s\n", len(copied), cap(copied), copied, copied)
	}

	// copy can be used to copy elements of slice to another slice

	// Syntax:
	// copy({destination slice}, {source slice})

	// Rules:
	// - type of both slices must be the same
	// - copy() only copies the elements based on the length of smallest slice
	//   if source slice has length of 5 and destination slice has length of 3
	//   only element 0-2 will be copied to destination slice
	// - copy() returns an integer which represents the number of copied elements

	source := []int{51, 53, 57, 59}

	// to copy the slice using 'copy()', first we need to create a destination slice
	dest := make([]int, len(source))

	// then perform the copy operation
	n := copy(dest, source)
	fmt.Println("copy()")
	fmt.Printf("'source' Length %d Capacity %d addr (backing array) %p %d\n", len(source), cap(source), source, source)
	fmt.Printf("'dest' Length %d Capacity %d addr (backing array) %p %d\n", len(dest), cap(dest), dest, dest)
	fmt.Printf("%d elements are copied\n", n)

	// though it is very convenient to copy using 'copy()', 'append()' is still a preferred way
	// to write a more concise code. Copying using 'copy()' requires the creation of destination slice
	// using 'make()' before the copy operation can be performed. Meanwhile, 'append()' can achieve
	// the same result by appending the source slice to a nil slice. The tradeoff is, a new backing
	// array is constantly created as the size grows

	dest2 := append([]int(nil), source...)
	fmt.Printf("'dest2' Length %d Capacity %d addr (backing array) %p %d\n", len(dest2), cap(dest2), dest2, dest2)

}
