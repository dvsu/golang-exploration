package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// get directory name from command line argument
	// and scan for subdirectories.
	// If found, write the subdirectories to a file.

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: go run main.go {dir1} {dir2} ... {dirN}")
		return
	}

	var output []byte

	for i, arg := range args {
		files, err := ioutil.ReadDir(arg)

		if err != nil {
			fmt.Println(err)
			return
		}

		output = append(output, fmt.Sprintf("%s\n", args[i])...)
		for _, file := range files {
			if file.IsDir() {
				output = append(output, fmt.Sprintf("    %s\n", file.Name())...)
			}
		}
		output = append(output, '\n')
	}

	err := ioutil.WriteFile("subdirs.txt", output, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
