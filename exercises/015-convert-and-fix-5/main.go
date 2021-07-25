package main

import "fmt"

func main() {
	// expected 1127
	min := int8(127)
	max := int16(1000)

	fmt.Println(max + int16(min))
}
