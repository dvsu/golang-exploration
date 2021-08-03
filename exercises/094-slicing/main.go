package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := "2 4 6 1 3 5"

	// process the data to generate the following output
	//     nums        : [2 4 6 1 3 5]
	//     evens       : [2 4 6]
	//     odds        : [1 3 5]
	//     middle      : [6 1]
	//     first 2     : [2 4]
	//     last 2      : [3 5]
	//     evens last 1: [6]
	//     odds last 2 : [3 5]

	var nums []int

	for _, n := range strings.Split(data, " ") {
		d, _ := strconv.Atoi(n)
		nums = append(nums, d)
	}

	evens := nums[:3]
	odds := nums[3:]

	fmt.Printf("nums         : %d\n", nums)
	fmt.Printf("evens        : %d\n", evens)
	fmt.Printf("odds         : %d\n", odds)
	fmt.Printf("middle       : %d\n", nums[2:4])
	fmt.Printf("first 2      : %d\n", nums[:2])
	fmt.Printf("last 2       : %d\n", nums[4:])
	fmt.Printf("evens last 1 : %d\n", evens[2:])
	fmt.Printf("odds last 2  : %d\n", odds[1:])
}
