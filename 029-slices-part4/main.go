package main

import "fmt"

func main() {
	// slice a slice
	colors := []string{
		"red",
		"green",
		"blue",
		"cyan",
		"yellow",
		"magenta",
		"black",
		"white",
		"grey",
		"brown",
		"purple",
		"pink",
		"violet",
		"orange",
	}

	fmt.Printf("%#v\n", colors[1:4]) // ["green", "blue", "cyan"]
	fmt.Printf("%#v\n", colors[:3])  // ["red", "green", "blue"]
	fmt.Printf("%#v\n", colors[5:])  // ["magenta", "black"]
	fmt.Printf("%#v\n", colors[:])   // ["red", "green", "blue", "cyan", "yellow", "magenta", "black"]

	// pagination using slice expressions
	// assuming 3 colors per page

	const pageSize = 3
	l := len(colors)

	for start := 0; start < l; start += pageSize {

		end := start + pageSize

		if end > l {
			end = l
		}

		fmt.Printf("Page %d: %s\n", (start/pageSize)+1, colors[start:end])
	}
}
