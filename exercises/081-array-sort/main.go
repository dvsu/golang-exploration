package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// take input of 5 numbers and sort it from minimum to maximum using bubblesort
	// it should be able to handle non-numeric inputs as well
	args := os.Args[1:]

	if len(args) < 1 || len(args) > 5 {
		fmt.Println("Usage: go run main.go [num_1] [num_2] ... [num_5]")
		return
	}

	var n [5]int
	var l float64
	var sum float64

	for i, e := range args {
		if d, err := strconv.Atoi(e); err != nil {
			n[i] = 0
			continue
		} else {
			n[i] = d
			l += 1
			sum += float64(d)
		}
	}

	fmt.Printf("Your numbers: %d\n", n)

	var (
		i int
		j int
	)

	max := len(n)
	// bubblesort
	for i = 0; i < max-1; i++ {
		for j = 0; j < max-i-1; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}

	fmt.Printf("Your sorted numbers: %d\n", n)
}
