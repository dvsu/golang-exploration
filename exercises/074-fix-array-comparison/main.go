package main

import "fmt"

func main() {
	// this section has to be fixed,
	// so the arrays are comparable

	// week := [...]string{"Monday", "Tuesday"}
	// wend := [4]string{"Saturday", "Sunday"}

	// fmt.Println(week != wend)

	// evens := [...]int{2, 4, 6, 8, 10}
	// evens2 := [...]int32{2, 4, 6, 8, 10}

	// fmt.Println(evens == evens2)

	// image := [5]byte{'h', 'i'}
	// var data [5]byte

	// fmt.Println(data == image)

	// expected output
	// true
	// true
	// false

	week := [4]string{"Monday", "Tuesday"}
	wend := [4]string{"Saturday", "Sunday"}

	fmt.Println(week != wend)

	evens := [...]int{2, 4, 6, 8, 10}
	evens2 := [...]int{2, 4, 6, 8, 10}

	fmt.Println(evens == evens2)

	// uint8 is 8 bits and byte is also 8 bits
	image := [5]uint8{'h', 'i'}
	var data [5]byte

	fmt.Println(data == image)
}
