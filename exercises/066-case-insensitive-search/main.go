package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Keyword missing. Please give a word to perform search")
		return
	}

	t := strings.Fields(`Judge Pepperleigh lived in a big house with 
hardwood floors and a wide piazza that 
looked over the lake from the top of Oneida Street.`)

queries:
	for _, q := range args {
	search:
		for _, word := range t {
			switch strings.ToLower(q) {
			case "and", "or", "the":
				break search
			}
			if strings.EqualFold(word, q) {
				fmt.Println(strings.ToLower(q))
				continue queries
			}
		}
	}
}
