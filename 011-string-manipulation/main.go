package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	l := len(msg)
	// repeat the underscore as many as l (length of message)
	signs := strings.Repeat("_", l)
	// convert the alphabets in the string to uppercase
	s := strings.ToUpper(signs + msg + signs)
	fmt.Println(s)
}
