package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// get numbers from command line and sort it
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please give some numbers")
		return
	}

	var nums []int

	for _, n := range args {
		if d, err := strconv.Atoi(n); err == nil {
			nums = append(nums, d)
		}
	}
	sort.Ints(nums)

	fmt.Println(nums)
}
