package main

import "fmt"

func main() {
	// conventional in-sequence number generation
	// const (
	// 	first = 1
	// 	second = 2
	// 	third = 3
	// 	fourth = 4
	// 	fifth = 5
	// )

	// it can be made simpler using iota. iota always starts at 0
	const (
		first = iota + 1
		second
		third
	)

	fmt.Println(first, second, third)

	// complex sequence can also be achieved using blank identifier to skip certain position
	const (
		eleventh = iota + 11
		twelfth
		_
		_
		_
		sixteenth
		_
		_
		_
		twentieth
	)

	fmt.Println(eleventh, twelfth, sixteenth, twentieth)
}
