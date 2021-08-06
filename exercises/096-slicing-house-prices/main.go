package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const (
		data = `Location,Size,Beds,Baths,Price
New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	args := os.Args[1:]

	rows := strings.Split(data, "\n")
	cols := strings.Split(rows[0], separator)

	start, end := 0, len(cols)

	l := len(args)

	// find the index position
	for i, e := range cols {
		if l >= 1 && e == args[0] {
			start = i
		}

		if l == 2 && e == args[1] {
			end = i + 1
		}
	}

	// reset the start to index 0 if the start index is larger than end
	if start > end {
		start = 0
	}

	// pretty print
	for _, row := range rows {

		r := strings.Split(row, separator)
		for x := start; x < end; x++ {
			fmt.Printf("%-15s", r[x])
		}
		fmt.Println()
	}
}
