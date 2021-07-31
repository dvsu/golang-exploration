package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// read input args
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Query word missing")
		return
	}

	// get path variables
	p := filepath.SplitList(os.Getenv("Path"))

	for _, q := range args {
		for _, e := range p {
			if strings.Contains(strings.ToLower(e), strings.ToLower(q)) {
				fmt.Println(e)
			}
		}
	}
}
