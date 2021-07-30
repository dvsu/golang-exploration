package main

import (
	"fmt"
	"os"
	"strconv"
)

func isLeapYear(y int) bool {
	if y%4 != 0 && (y%100 == 0 || y%400 != 0) {
		return false
	} else {
		return true
	}
}

func main() {
	if args := os.Args; len(args) != 2 {
		fmt.Println("Give me a year number")
	} else if y, err := strconv.Atoi(args[1]); err != nil {
		fmt.Printf("%q is not a valid year.\n", args[1])
	} else {
		if isLeapYear(y) {
			fmt.Printf("%d is a leap year.\n", y)
		} else {
			fmt.Printf("%d is not a leap year.\n", y)
		}
	}
}
