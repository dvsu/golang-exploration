package main

import "fmt"

func main() {
	const (
		_ = iota * 3
		Spring
		Summer
		Fall
		Winter
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}
