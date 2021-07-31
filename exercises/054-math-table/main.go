package main

import (
	"fmt"
	"os"
	"strconv"
)

func addition(s int) {
	for m := 0; m <= s; m++ {
		fmt.Printf("%8d", m)
		for n := 0; n <= s; n++ {
			fmt.Printf("%8d", m+n)
		}
		fmt.Println("")
	}
}

func substraction(s int) {
	for m := 0; m <= s; m++ {
		fmt.Printf("%8d", m)
		for n := 0; n <= s; n++ {
			fmt.Printf("%8d", m-n)
		}
		fmt.Println("")
	}
}

func multiplication(s int) {
	for m := 0; m <= s; m++ {
		fmt.Printf("%8d", m)
		for n := 0; n <= s; n++ {
			fmt.Printf("%8d", m*n)
		}
		fmt.Println("")
	}
}

func division(s int) {
	for m := 0; m <= s; m++ {
		fmt.Printf("%8d", m)
		for n := 0; n <= s; n++ {
			if n == 0 {
				fmt.Printf("%8s", "-")
				continue
			}
			fmt.Printf("%8.2f", float64(m)/float64(n))
		}
		fmt.Println("")
	}
}

func main() {

	args := os.Args

	var (
		s   int
		err error
	)

	l := len(args)

	switch {
	case l == 2:
		fmt.Printf("Size is missing\n")
		fallthrough
	case l > 3:
		fallthrough
	case l == 1:
		fmt.Printf("Usage: [op=*/+-] [size]\n")
		return
	}

	if s, err = strconv.Atoi(args[2]); err != nil {
		fmt.Println("Wrong table size")
		return
	}

	// print header
	fmt.Printf("%8s", args[1])
	for i := 0; i <= s; i++ {
		fmt.Printf("%8d", i)
	}
	fmt.Println("")

	// print content
	switch args[1] {
	case "*":
		multiplication(s)
	case "/":
		division(s)
	case "+":
		addition(s)
	case "-":
		substraction(s)
	}
}
