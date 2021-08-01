package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// search for book title that contains queried word
	books := [...]string{
		"Chronicles of Avonlea",
		"Rainbow Valley",
		"Anne's House of Dreams",
		"Rilla of Ingleside",
		"The Story of My Life",
		"The Story Girl",
		"Sunshine Sketches of a Little Town",
		"The Golden Road",
		"Anne of Green Gables",
		"Anne of Avonlea",
		"Kilmeny of the Orchard",
	}

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Usage: go run main.go [queried word]")
		return
	}

	var count int

	fmt.Printf("Search results for %q:\n", args[0])
	for _, b := range books {
		if strings.Contains(strings.ToLower(b), strings.ToLower(args[0])) {
			fmt.Printf("+ %s\n", b)
			count += 1
		}
	}

	switch {
	case count == 0:
		fmt.Println("No books found")
	case count == 1:
		fmt.Printf("%d book found\n", count)
	default:
		fmt.Printf("%d books found\n", count)
	}
}
