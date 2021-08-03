package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// load data to slices, calculate the average, and pretty print it
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		locations []string
		sizes     []int
		beds      []int
		baths     []int
		prices    []int
	)

	headers := strings.Split(header, separator)

	rows := strings.Split(data, "\n")

	for _, n := range headers {
		fmt.Printf("%-15s", n)
	}
	fmt.Println()
	fmt.Printf("%s\n", strings.Repeat("=", 70))

	for _, row := range rows {
		r := strings.Split(row, separator)
		locations = append(locations, r[0])
		if s, err := strconv.Atoi(r[1]); err != nil {
			sizes = append(sizes, 0)
		} else {
			sizes = append(sizes, s)
		}

		if be, err := strconv.Atoi(r[2]); err != nil {
			beds = append(beds, 0)
		} else {
			beds = append(beds, be)
		}

		if ba, err := strconv.Atoi(r[3]); err != nil {
			baths = append(baths, 0)
		} else {
			baths = append(baths, ba)
		}

		if p, err := strconv.Atoi(r[4]); err != nil {
			prices = append(prices, 0)
		} else {
			prices = append(prices, p)
		}
	}

	var (
		sumSizes float64
		sumBeds  float64
		sumBaths float64
		sumPrice float64
	)

	for i := range rows {
		fmt.Printf("%-15s%-15d%-15d%-15d%-15d\n", locations[i], sizes[i], beds[i], baths[i], prices[i])
		sumSizes += float64(sizes[i])
		sumBeds += float64(beds[i])
		sumBaths += float64(baths[i])
		sumPrice += float64(prices[i])
	}

	fmt.Printf("%s\n", strings.Repeat("=", 70))

	// calculate and print the average
	fmt.Printf("%-15s", "")
	fmt.Printf("%-15.2f", sumSizes/float64(len(sizes)))
	fmt.Printf("%-15.2f", sumBeds/float64(len(beds)))
	fmt.Printf("%-15.2f", sumBaths/float64(len(baths)))
	fmt.Printf("%-15.2f", sumPrice/float64(len(prices)))
}
