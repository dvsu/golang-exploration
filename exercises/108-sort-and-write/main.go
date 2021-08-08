package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func main() {
	// sort string input from command line and write
	// the result to a file
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: go run main.go {string1} {string2} ... {stringN}")
		return
	}

	sort.Strings(args)

	var size int
	// preallocate memory for backing array
	for _, arg := range args {
		size += len(arg) + 1 // include newline character
	}

	s := make([]byte, 0, size)

	for _, arg := range args {
		s = append(s, arg...)
		s = append(s, '\n')
	}

	// write the output to file
	err := ioutil.WriteFile("sorted.txt", s, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
