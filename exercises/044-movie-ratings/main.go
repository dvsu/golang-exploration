package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if args := os.Args; len(args) != 2 {
		fmt.Println("Requires age")
	} else if a, err := strconv.Atoi(args[1]); err != nil || a < 0 {
		fmt.Printf("Wrong age %q\n", args[1])
	} else {
		if a > 17 {
			fmt.Println("R-Rated")
		} else if a > 12 {
			fmt.Println("PG-13")
		} else {
			fmt.Println("PG-Rated")
		}
	}
}
