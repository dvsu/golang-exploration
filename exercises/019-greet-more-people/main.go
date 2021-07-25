package main

import (
	"fmt"
	"os"
)

func main() {
	// go run main.go aaa bbb ccc
	l := len(os.Args) - 1
	fmt.Printf("There are %v people!\n", l)
	for c, args := range os.Args {
		if c > 0 {
			fmt.Printf("Hello great %v!\n", args)
		}
	}
	fmt.Println("Nice to meet you all.")
}
