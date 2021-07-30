package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if args := os.Args; len(args) != 2 {
		fmt.Println("Pick a number")
	} else if n, err := strconv.Atoi(args[1]); err != nil {
		fmt.Printf("%q is not a number\n", args[1])
	} else {
		if n%2 == 0 && n%8 == 0 {
			fmt.Printf("%d is an even number and it's divisible by 8\n", n)
		} else if n%2 == 0 {
			fmt.Printf("%d is an even number\n", n)
		} else {
			fmt.Printf("%d is an odd number\n", n)
		}
	}
}
