package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func isLeapYear(y int) bool {
	if y%4 != 0 && (y%100 == 0 || y%400 != 0) {
		return false
	} else {
		return true
	}
}

func main() {

	var (
		args []string
		m    string
		d    int = 28 // default value of non-leap February
	)
	if args = os.Args; len(args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	leap := isLeapYear(time.Now().Year())

	m = strings.ToLower(args[1])

	switch m {
	case "january", "march", "may", "july", "august", "october", "december":
		d = 31
	case "april", "june", "september", "november":
		d = 30
	case "february":
		if leap {
			d = 29
		}
	default:
		fmt.Printf("%q is not a month.\n", args[1])
		return
	}

	fmt.Printf("%q has %d days.\n", args[1], d)
}
