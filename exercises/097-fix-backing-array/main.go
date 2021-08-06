package main

import "fmt"

func main() {
	// (INCORRECT)
	//  Mine         : [-50 -100 -150 25 30 50]
	//  Original nums: [-50 -100 -150]
	//
	//
	// EXPECTED OUTPUT
	//
	//  Mine         : [-50 -100 -150]
	//  Original nums: [56 89 15]

	nums := []int{56, 89, 15, 25, 30, 50}

	// BEFORE
	// mine := num

	mine := append([]int{}, nums[:3]...)
	// alternative
	// mine := append([]int(nil), nums[:3]...)

	mine[0], mine[1], mine[2] = -50, -100, -150

	fmt.Println("Mine         :", mine)
	fmt.Println("Original nums:", nums[:3])
}
