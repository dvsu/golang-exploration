package main

import "fmt"

func main() {
	fmt.Println("First line")
	foo()
	fmt.Println("Additional line")
	for i := 0; i < 30; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("Last one")
}

func foo() {
	fmt.Println("Inside foo")
}
