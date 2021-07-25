package main

import (
	"fmt"
	"os"
)

func main() {
	// go run main.go {name}
	// expected
	//    Hi {name}
	//    How are you?
	fmt.Printf("Hi %v\nHow are you?", os.Args[1])
}
