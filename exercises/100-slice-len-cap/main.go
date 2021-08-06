package main

import "fmt"

func main() {
	// part 1
	fmt.Println("Part 1")
	// create new nil slice
	// var games []string
	// fmt.Printf("Length: %d Capacity: %d\n", len(games), cap(games))

	// append new elements
	// games = []string{"pacman", "mario", "tetris", "doom"}
	// fmt.Printf("Length: %d Capacity: %d\n", len(games), cap(games))

	// redeclare slice using slice literal
	games := []string{"pacman", "mario", "tetris", "doom"}

	// part 2
	fmt.Println("Part 2")
	// loop from 0 to 4, slice the 'games' slice, element by element
	// and print its length and capacity
	for i := 0; i <= 4; i++ {
		games = games[:i]
		fmt.Printf("Length: %d Capacity: %d Values: %s\n", len(games), cap(games), games)
	}

	// part 3
	fmt.Println("Part 3")
	// slice the 'games' slice up to zero element
	zero := games[:0]
	fmt.Printf("'games' backing array addr %p\n", games)
	fmt.Printf("'zero' backing array addr %p\n", zero)

	// check length and capacity of 'games' slice
	fmt.Printf("'games' Length: %d Capacity: %d Values: %s\n", len(games), cap(games), games)
	fmt.Printf("'zero' before | Length: %d Capacity: %d Values: %s\n", len(zero), cap(zero), zero)

	// append a new element to 'zero' slice
	zero = append(zero, "snake")
	fmt.Printf("'zero' after | Length: %d Capacity: %d Values: %s\n", len(zero), cap(zero), zero)

	// append 5 elements using loop
	for _, e := range []string{"zelda", "mega man", "sonic", "contra", "super mario bros"} {
		zero = append(zero, e)
		fmt.Printf("'zero' after | Length: %d Capacity: %d Values: %q\n", len(zero), cap(zero), zero)
	}

	// part 4
	fmt.Println("Part 4")
	// use for range loop to slice 'zero' slice element by element
	for i := range zero {
		slc := zero[:i]
		fmt.Printf("'zero' after | Length: %d Capacity: %d Values: %q\n", len(slc), cap(slc), slc)
	}

	// part 5
	fmt.Println("Part 5")
	// slice 'zero' slice up to its capacity
	zero = zero[:cap(zero)]
	for i := range zero {
		slc2 := zero[:i]
		fmt.Printf("'zero' after | Length: %d Capacity: %d Values: %q\n", len(slc2), cap(slc2), slc2)
	}

	// part 6
	fmt.Println("Part 6")
	// change an element of 'games' and 'zero' slices and prove both backing array is not same
	games[0] = "minecraft"
	zero[0] = "final fantasy"
	fmt.Printf("'games' final | Length: %d Capacity: %d Values: %q %p\n", len(games), cap(games), games, games)
	fmt.Printf("'zero' final | Length: %d Capacity: %d Values: %q %p\n", len(zero), cap(zero), zero, zero)
}
