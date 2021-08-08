package main

import (
	"fmt"
	"strings"
)

// CURRENT OUTPUT
//
//  yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday
//
// EXPECTED OUTPUT
//
//  yesterday all my troubles seemed so far away
//  now it looks as though they are here to stay
//  oh I believe in yesterday
//
// RESTRICTIONS
//
//  + Don't use `append()`, use `copy()` instead.
//
//  + Don't cheat like this:
//
//    fmt.Println(lyric[:8])
//    fmt.Println(lyric[8:18])
//    fmt.Println(lyric[18:23])
//
//  + Create a new slice that contains the sentences
//    with line endings.

func main() {
	lyric := strings.Fields(`yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday`)

	// solution 1 starts here
	fmt.Println("Solution 1")
	// create 2d slice with 3 rows
	fix := make([][]string, 3)

	// mark the index of new sentences
	start := [...]int{0, 8, 18, len(lyric)}

	// iterate through the rows and copy the elements from 'lyric' to 'fix' columns, row by row
	for i := range fix {
		fix[i] = make([]string, len(lyric[start[i]:start[i+1]]))
		copy(fix[i], lyric[start[i]:start[i+1]])
	}

	for _, row := range fix {
		for j, w := range row {
			fmt.Print(w)
			if j != len(row)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	// solution 1 ends here

	// solution 2 starts here
	fmt.Println("Solution 2")
	// use 1d slice and insert newline character "\n" at every end of sentence

	fix2 := make([]string, len(lyric)+3)

	for i := 0; i < 3; i++ {
		copy(fix2[start[i]+i:start[i+1]+i], lyric[start[i]:start[i+1]])
		fix2[start[i+1]+i] = "\n"
	}

	for _, w := range fix2 {
		fmt.Print(w)
		if w != "\n" {
			fmt.Print(" ")
		}
	}
	// solution 2 ends here

	// solution 3 starts here
	fmt.Println("Solution 3")
	// use 1d slice and insert newline character "\n" at every end of sentence
	// the number of elements copied will be used as helper for cutting index

	// the number of words per line
	end := []int{8, 10, 5}

	fix3 := make([]string, len(lyric)+3)

	for i, n := 0, 0; n < len(lyric); i++ {
		n += copy(fix3[n+i:], lyric[n:n+end[i]])
		fix3[n+i] = "\n"
	}

	for _, w := range fix3 {
		fmt.Print(w)
		if w != "\n" {
			fmt.Print(" ")
		}
	}
	// solution 3 ends here
}
