package main

import (
	"fmt"
	"os"
)

func main() {
	// expected "There are 3 names."
	count := len(os.Args) - 1
	fmt.Printf("There are %d names.\n", count)
}
