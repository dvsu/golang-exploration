package main

import (
	"fmt"
	"sort"
)

// avoid package level variable if possible.
// Because it is accessible from all functions in the same package,
// it may introduce unwanted bugs.
// var x int // bad idea

func main() {
	fmt.Println(filterDataSorted([]int{20, 30, 12, 15, 19, 1, 11, 21, 9}, 15))
	fmt.Println(filterData([]int{20, 30, 12, 15, 19, 1, 11, 21, 9}, 15))
}

// typical function in Go
func filterDataSorted(data []int, lim int)  []int {
	
	var sorted []int

	for _, n := range data {
		if n < lim {
			sorted = append(sorted, n)
		}
	}

	sort.Ints(sorted)

	return sorted
}

// we can also name the return variable
// this return variable is equal to
// var filtered []int
// the return can be intentionally left blank
// because Go knows the return variable is 'filtered'
func filterData(data []int, lim int) (filtered []int) {
	for _, n := range data {
		if n < lim {
			filtered = append(filtered, n)
		}
	}
	return
}