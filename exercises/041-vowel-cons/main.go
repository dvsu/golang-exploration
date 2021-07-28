package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	l := len(os.Args) - 1
	vowel := "aeiou"

	if l != 1 || len(os.Args[1]) != 1 {
		fmt.Println("No or too many arguments. Give me an alphabet")
		return
	}

	c := os.Args[1]

	if strings.ContainsAny(vowel, c) {
		fmt.Printf("Alphabet %q is a vowel\n", c)
	} else {
		fmt.Printf("Alphabet %q is a consonant\n", c)
	}
}
