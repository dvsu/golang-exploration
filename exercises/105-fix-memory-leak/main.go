package main

import (
	"fmt"
	"io/ioutil"

	"github.com/dvsu/golang-exploration/exercises/105-fix-memory-leak/api"
)

//  CURRENT OUTPUT
//
//    > Memory Usage: 113 KB (before)
//    > Memory Usage: 65651 KB (after)

//  EXPECTED OUTPUT
//
//    > Memory Usage: 116 KB (before)
//    > Memory Usage: 118 KB (after)

func main() {
	// reports the initial memory usage
	api.Report()

	// returns a slice with 10 million elements.
	// it allocates 65 MB of memory space.
	millions := api.Read()

	// let's print the original pointer
	fmt.Printf("'millions' Original pointer to backing array: %p\n", millions)

	// Before
	// last10 := millions[len(millions)-10:]

	/*
		// After: Solution 1 (with explanation)
		// create a slice with 10 empty elements
		last10 := make([]int, 10)

		// copy the last 10 elements from 'millions' slice to 'last10'
		copy(last10, millions[len(millions)-10:])

		// At this point, 'last10' and 'millions' have different backing arrays
		// As 'millions' slice is not needed anymore, simply reassign the 'pointer'
		// to the new backing array belongs to 'last10'
		millions = last10

		// After reassignment, both will have the same backing array
		// it also means 'millions' will lose its pointer to the old
		// backing array. Then, Go will claim the unused memory
		// by performing garbage collection behind the scene.
		// Therefore, the original size of 60MB+ has shrunk to
		// the size of 10-element backing array (1xx KB)
		fmt.Printf("\nLast 10 elements: %d\n\n", last10)

		// let's print the pointer to backing array again and compare the
		// address to the previous one.
		fmt.Printf("'millions' New pointer to backing array: %p last10: %p\n", millions, last10)
	*/

	// After: Solution 2 (alternative one-line solution with explanation)
	// because 'last10' is only a temporary slice to store the last 10 elements
	// the process can be done directly to the original slice by
	// appending the last 10 elements to nil slice and reassigning it to the original slice
	millions = append([]int(nil), millions[len(millions)-10:]...)
	fmt.Printf("'millions' Updated pointer to backing array: %p\n", millions)
	api.Report()

	fmt.Fprintln(ioutil.Discard, millions[0])
}
