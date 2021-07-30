package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		args []string
		f    float64
		err  error
	)
	if args = os.Args; len(args) != 2 {
		fmt.Println("Please provide the magnitude of the earthquake")
		return
	}

	if f, err = strconv.ParseFloat(args[1], 64); err != nil {
		fmt.Println("Unknown magnitude")
		return
	}

	switch {
	case f >= 10.0:
		fmt.Println("massive")
	case f >= 8.0:
		fmt.Println("great")
	case f >= 7.0:
		fmt.Println("major")
	case f >= 6.0:
		fmt.Println("strong")
	case f >= 5.0:
		fmt.Println("moderate")
	case f >= 4.0:
		fmt.Println("light")
	case f >= 3.0:
		fmt.Println("minor")
	case f >= 2.0:
		fmt.Println("very minor")
	default:
		fmt.Println("micro")
	}

}
