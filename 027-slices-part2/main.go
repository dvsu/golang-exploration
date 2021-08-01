package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	// populate uniques numbers using slices to
	// show its superiority over array

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Usage: go run main.go [n-unique-numbers]")
		return
	}

	max, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("Usage: go run main.go [n-unique-numbers]")
		return
	}

	var uniques []int

	rand.Seed(time.Now().UnixNano())

loop:
	for len(uniques) < max {
		r := rand.Intn(max) + 1

		for _, n := range uniques {
			if n == r {
				continue loop
			}
		}

		// if it manages to complete the for loop,
		// it means the number is unique
		// then, assign the r value as an element
		// of 'uniques' slice
		uniques = append(uniques, r)
	}
	fmt.Printf("Unique numbers: %#v\n", uniques)

	// Go has many built-in functions to process slices
	// which cannot be done on array

	// for example, sorting a slice
	sort.Ints(uniques)
	fmt.Printf("Sorted unique numbers: %#v\n", uniques)

	// this built-in sort cannot be done on an array
	arr := [10]int{6, 8, 3, 21, 7, 1, 6, 9, 3, 4}

	// this is not allowed
	// sort.Ints(arr)

	// however, it can be done by converting type [10]int to []int first
	// because the sort.Ints() function only accepts []int as input

	// convert array to slice
	sort.Ints(arr[:])

	// now array 'arr' is sorted
	fmt.Printf("Sorted array: %#v\n", arr)
}
