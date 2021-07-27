package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// len() is a built-in function. import is not required
	// len returns the number of bytes
	word1 := "Computer"

	// it will print 8, as standard English character is equal to 1 byte
	fmt.Println(len(word1))

	word2 := "Buenos días"

	// it will print 12, as non-standard English character, í, is equal to 2 bytes
	// whitespace is equal to 1 byte
	fmt.Println(len(word2))

	// the solution to the problem is by using RuneCountInString function from utf8 package
	// print statement below will output 11
	fmt.Println(utf8.RuneCountInString(word2))
}
