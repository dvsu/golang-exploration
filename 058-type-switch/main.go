package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Type switch can be used to detect and extract
	// the dynamic value from an interface value
	// It works like a type assertion, but is
	// done inside switch statement.

	/*
						Syntax

						        interface
				                  value
		  extracted value -----.    |
						       V    V
						switch e := v.(type) {

						}

						The type of 'e' will change, depending on
						the extracted value

	*/

	var source interface{}
	
	source = 12345
	fmt.Printf("%q, type %[1]T\n", format(source))
	source = "another one"
	fmt.Printf("%q, type %[1]T\n", format(source))
	source = []int{10, 20, 30}
	fmt.Printf("%q, type %[1]T\n", format(source))
}

// a function that will read various dynamic data types from
// en empty interface, convert it, and return as string
func format(v interface{}) string {
	var t int

	switch e := v.(type) {
	case int:
		t = e
		return strconv.Itoa(t)
	case string:
		return e
	default:
		return ""
	}
}