package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var a int

	for i := 0; i < 5; i++ {
		a += i
	}
	fmt.Println(a)

	// there is no while loop in Go.
	// However, mimicking while loop is as simple as using the for loop without any statement
	// code block below is equivalent to while true {...}
	j := 0
	b := 0

	for {
		if j > 3 {
			break
		}
		b += j
		j++
	}

	fmt.Println(b)

	// create a multiplication table by using nested for-loop
	args := os.Args
	var (
		max int
		err error
	)

	if len(args) != 2 {
		fmt.Println("Please input the table size")
		return
	}

	if max, err = strconv.Atoi(args[1]); err != nil {
		fmt.Println("Please input the table size as a number")
		return
	}

	// print header
	fmt.Printf("%s\n", strings.Repeat("=", 90))
	fmt.Printf("%8s", "mul")
	for i := 1; i <= max; i++ {
		fmt.Printf("%8d", i)
	}
	fmt.Println("")

	for i := 1; i <= max; i++ {
		fmt.Printf("%8d", i)
		for j := 1; j <= max; j++ {
			fmt.Printf("%8d", i*j)
		}
		fmt.Println("")
	}

	fmt.Printf("%s\n", strings.Repeat("=", 90))
	fmt.Printf("%8s", "div")
	for i := 1; i <= max; i++ {
		fmt.Printf("%8d", i)
	}
	fmt.Println("")

	// create a division table using nested for-loop
	for i := 1; i <= max; i++ {
		fmt.Printf("%8d", i)
		for j := 1; j <= max; j++ {
			fmt.Printf("%8.3f", float64(j)/float64(i))
		}
		fmt.Println("")
	}

}
