package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		args  []string
		scale string
		r     string
	)

	if args = os.Args; len(args) != 2 {
		fmt.Println("Please provide the scale of the earthquake in Richter Scale")
		return
	}

	scale = strings.ToLower(args[1])

	switch scale {
	case "massive":
		r = "10+"
	case "great":
		r = "8.0 - 8.9"
	case "major":
		r = "7.0 - 7.9"
	case "strong":
		r = "6.0 - 6.9"
	case "moderate":
		r = "5.0 - 5.9"
	case "light":
		r = "4.0 - 4.9"
	case "minor":
		r = "3.0 - 3.9"
	case "very minor":
		r = "2.0 - 2.9"
	case "micro":
		r = "1.0 - 1.9"
	default:
		r = "unknown"
	}

	fmt.Printf("%s's Ritcher Scale is %s\n", scale, r)
}
