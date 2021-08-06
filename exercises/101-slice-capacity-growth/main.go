package main

import "fmt"

func main() {
	// create nil slice and loop 1e7 times
	// append an element for each iteration
	// print when capacity changes

	var (
		nums []int
		c    int
	)

	for i := 0; i < 1e7; i++ {
		if c == 0 || cap(nums) != c {
			fmt.Printf("len:%-10d cap:%-10d growth:%.3f\n", len(nums), cap(nums), float64(cap(nums))/float64(c))
			c = cap(nums)
		}
		nums = append(nums, 0)
	}
}
