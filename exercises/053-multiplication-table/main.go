package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	var (
		s   int
		err error
	)

	if len(args) != 2 {
		fmt.Println("Give me the size of the table")
		return
	}

	if s, err = strconv.Atoi(args[1]); err != nil {
		fmt.Println("Wrong size")
		return
	}

	// print header
	fmt.Printf("%8s", "X")
	for i := 0; i <= s; i++ {
		fmt.Printf("%8d", i)
	}
	fmt.Println("")

	// print contents
	for m := 0; m <= s; m++ {
		fmt.Printf("%8d", m)
		for n := 0; n <= s; n++ {
			fmt.Printf("%8d", n*m)
		}
		fmt.Println("")
	}
}
