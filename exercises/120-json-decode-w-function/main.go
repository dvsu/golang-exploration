package main

import (
	"bufio"
	"os"
)

func main() {
	// Refactor exercise 119 by using functions

	games := load()
	var (
		content []game
		cont bool
	)
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {

		input := in.Bytes()

		content, cont = runCmd(string(input), games)
		if !cont {
			return
		}
		if content != nil {
			games = content
		}	
	}
}
