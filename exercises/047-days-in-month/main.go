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

	if m = strings.ToLower(args[1]); m == "january" ||
		m == "march" ||
		m == "may" ||
		m == "july" ||
		m == "august" ||
		m == "october" ||
		m == "december" {
		d = 31
	} else if m == "april" ||
		m == "june" ||
		m == "september" ||
		m == "november" {
		d = 30
	} else if m == "february" {
		if leap {
			d = 29
		}
	} else {
		fmt.Printf("%q is not a month.\n", args[1])
		return
	}

	fmt.Printf("%q has %d days.\n", args[1], d)
}
