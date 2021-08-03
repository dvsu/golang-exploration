package main

import "fmt"

func main() {
	// add a new element to a slice using append
	slc0 := []int{99, 100}

	slc0 = append(slc0, 101)

	// output [99, 100, 101]
	fmt.Println(slc0)

	// adding multiple elements is also possible
	slc1 := []int{9, 0, 1}

	slc1 = append(slc1, 2, 3, 4, 5)

	// output [9, 0, 1, 2, 3, 4, 5]
	fmt.Println(slc1)

	// append slice to another slice using ellipsis operator
	slc2 := []int{1, 3, 5}
	slc3 := []int{2, 6, 8}

	slc3 = append(slc3, slc2...)

	// output [2, 6, 8, 1, 3, 5]
	fmt.Println(slc3)

}
