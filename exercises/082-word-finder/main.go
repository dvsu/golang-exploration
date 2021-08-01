package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = "Usage: go run main.go [queried_word_1] [queried_word_2] ... [queried_word_n]"

func main() {
	t := strings.Fields(`The honey-tinted autumn sunshine was 
falling thickly over the crimson and amber maples around old Abel Blair’s door`)

	// find queried words in the given text
	// it should be able to filter non essential words, such as
	// and. the, or, is, are, was, were
	// the non essential words have to be stored in an array

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println(usage)
		return
	}

	filter := [...]string{
		"and",
		"the",
		"or",
		"is",
		"are",
		"was",
		"were",
	}

queries:
	for _, q := range args {

		for _, f := range filter {
			if strings.EqualFold(q, f) {
				continue queries
			}
		}

		for j, w := range t {
			if strings.EqualFold(q, w) {
				fmt.Printf("#%2d: %q\n", j+1, w)
				continue queries
			}
		}
	}
}
