package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// comparing 2 slices
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	// split to convert string into slice
	slcA := strings.Split(namesA, ",")

	// remove any trailing whitespace
	for i, n := range slcA {
		slcA[i] = strings.TrimSpace(n)
	}

	// sort both slices
	sort.Strings(slcA)
	sort.Strings(namesB)

	// compare each element
	for i, n := range slcA {
		if n != namesB[i] {
			fmt.Println("Not equal")
			return
		}
	}
	fmt.Println("Equal")
}
