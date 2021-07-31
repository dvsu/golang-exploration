package main

import "fmt"

func main() {
	// sum numbers from 1 to 10
	t := 0
	for i := 1; i <= 10; i++ {
		t += i
		fmt.Printf("%d ", i)
		if i != 10 {
			fmt.Printf("+ ")
		}
	}
	fmt.Printf("= %d\n", t)
}
