package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// example of string-to-integer conversion and store it in 2d slice
	data := `10 20 30 40 50
60 70 80 90 100
110 120 130 140 150
160 170 180 190 200
210 220 230 240 250`

	rows := strings.Split(data, "\n")

	// create empty 2d slice with 5 rows
	matrix := make([][]int, len(rows))

	for i, row := range rows {
		// split each row into multiple elements of string
		srow := strings.Fields(row)

		// create a row with 5 columns
		matrix[i] = make([]int, len(srow))

		// convert each element to int and store it in 2d array
		for j, s := range srow {
			n, _ := strconv.Atoi(s)
			matrix[i][j] = n
		}
	}

	fmt.Printf("%d\n", matrix)

	// example of 2d slice math operation

	// row sum
	for i, row := range matrix {
		var sum int

		for _, n := range row {
			sum += n
		}

		fmt.Printf("Row %d: %d\n", i, sum)
	}

	// column sum
	for col := 0; col < len(matrix[0]); col++ {
		var sum int
		for row := 0; row < len(matrix); row++ {
			sum += matrix[row][col]
		}
		fmt.Printf("Column %d: %d\n", col, sum)
	}
}
