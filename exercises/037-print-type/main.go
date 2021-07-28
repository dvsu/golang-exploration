package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Type of %d is %[1]T\n", 27)
	fmt.Printf("Type of %.8f is %[1]T\n", 2.718281828459045)
	fmt.Printf("Type of %s is %[1]T\n", "computer")
	fmt.Printf("Type of %t is %[1]T\n", false)

	// pass two arguments from command line
	// go run main.go abc def
	fmt.Printf("First and second arguments are %s and %s\n", os.Args[1], os.Args[2])
}
