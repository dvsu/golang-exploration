package main

import (
	"fmt"
	"sort"
)

func main() {
	// sort 3 elements in the middle of array (odd-length)

	// EXPECTED OUTPUT
	//
	//  Original: [pacman mario tetris doom galaga frogger asteroids simcity metroid defender rayman tempest ultima]
	//  Sorted  : [pacman mario tetris doom galaga asteroids frogger simcity metroid defender rayman tempest ultima]

	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	// odd by default
	mid := len(items) / 2 // index starts from 0
	start, end := mid-1, mid+2
	// if even, change the mid, start, and end values
	if len(items)%2 == 0 {
		start = mid - 2
	}

	fmt.Println("Original:", items)
	sorted := items[start:end]
	sort.Strings(sorted)
	fmt.Println("Sorted  :", items)
}
