package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	usage          = "Usage: go run main.go [start] [optional: end]"
	reversedOrder  = "'end' must be larger than 'start'"
	negativeNumber = "'start' and 'end' values must be positive"
)

func main() {
	// use args input from command line to perform slicing

	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}

	args := os.Args[1:]

	var (
		start  int
		end    int
		single bool
		sErr   error
		eErr   error
	)

	switch {
	case len(args) > 2:
		fallthrough
	case len(args) == 0:
		fmt.Println(usage)
		return
	case len(args) == 1:
		if start, sErr = strconv.Atoi(args[0]); sErr != nil {
			fmt.Println(usage)
			return
		}
		single = true
	case len(args) == 2:
		start, sErr = strconv.Atoi(args[0])
		end, eErr = strconv.Atoi(args[1])
		if sErr != nil || eErr != nil {
			fmt.Println(usage)
			return
		}
		single = false
	}

	switch {
	case !single && start > end:
		fmt.Println(reversedOrder)
		return
	case start < 0:
		fallthrough
	case end < 0:
		fmt.Println(negativeNumber)
		return
	}

	if single {
		fmt.Printf("%s\n", ships[start:])
	} else {
		fmt.Printf("%s\n", ships[start:end])
	}

}
