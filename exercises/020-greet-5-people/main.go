package main

import (
	"fmt"
	"os"
)

func main() {
	//go run main.go aaa bbb ccc ddd eee
	// restriction: do not use variables
	fmt.Printf("There are %d people!\n", len(os.Args)-1)
	for c, arg := range os.Args {
		if c > 0 {
			fmt.Printf("Hello great %v!\n", arg)
		}
	}
	fmt.Println("Nice to meet you all.")
}
