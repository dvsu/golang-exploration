package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func main() {
	// sort string input from command line and write
	// the result to a file, similar to part 1
	// In addition, the output also includes
	// index, e.g.
	// 1. eight
	// 2. three
	// 3. two

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: go run main.go {string1} {string2} ... {stringN}")
		return
	}

	sort.Strings(args)

	var size int
	// preallocate memory for backing array
	for _, arg := range args {
		size += len(arg) + 4 // include newline character, index number, dot, and whitespace
	}

	s := make([]byte, 0, size)

	for i, arg := range args {
		s = strconv.AppendInt(s, int64(i+1), 10)
		s = append(s, fmt.Sprintf(". %s\n", arg)...)
	}

	// write the output to file
	err := ioutil.WriteFile("sorted.txt", s, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
