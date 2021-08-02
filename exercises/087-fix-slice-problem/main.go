package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	planets := []string{
		"Mercury",
		"Venus",
		"Earth",
	}

	clouds := []string{
		"DigitalOcean",
		"AWS",
		"IBM Cloud",
		"Azure",
		"GCP",
	}

	sort.Strings(clouds)

	nums := [...]int{8, 5, 3, 2, 9, 0, 12, -99}

	// use slice expression to convert array to slice, then sort it
	sort.Ints(nums[:])

	//  strings.Join() to convert slice of string to string
	fmt.Printf("%q\n", strings.Join(planets, " + "))

	fmt.Printf("%q\n", clouds)
	fmt.Printf("%d\n", nums)
}
