package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// calculate the average of maximum 5 numbers
	// it should be able to handle unknown input, e.g. non-numeric
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

	switch l {
	case 0:
		fmt.Println("Average: 0")
	default:
		fmt.Printf("Average: %.2f\n", sum/l)
	}
}
