package main

import (
	"fmt"
	"strings"
)

func main() {
	// using multidimensional array and pretty print the elements

	data := [...][3]string{
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	// print header
	fmt.Printf("%-12s%-12s%-12s\n", "First Name", "Last Name", "Nickname")
	fmt.Printf("%s\n", strings.Repeat("=", 40))

	// print content
	for _, p := range data {
		for _, e := range p {
			fmt.Printf("%-12s", e)
		}
		fmt.Println()
	}
}
