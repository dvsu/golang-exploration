package main

import "fmt"

func main() {

	// multidimensional array
	// create 3x5 float64 array

	scoreboard := [...][5]float64{
		{4, 5, 3, 2, 7},
		{1, 6, 3, 8, 4},
		{9, 2, 7, 2, 5},
	}

	const N = float64(len(scoreboard) * len(scoreboard[0]))

	// find the average score
	var sum float64

	for _, scores := range scoreboard {
		for _, score := range scores {
			sum += score
		}
	}

	fmt.Printf("Average score: %.2f\n", sum/N)
}
