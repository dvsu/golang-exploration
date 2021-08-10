package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, word := range words {
		fmt.Printf("%q\n", word)
		// print the byte and rune length of the strings
		fmt.Printf("  Byte-length: %d bytes | Length: %d characters\n", len(word), utf8.RuneCountInString(word))

		// Print bytes of string and runes of string in hex
		fmt.Printf("%sBytes of string: % -12x\n", strings.Repeat(" ", 2), word)
		for _, r := range word {
			fmt.Printf("%sRune '%c' hex value:% [2]x\n", strings.Repeat(" ", 4), r)
		}

		// print the first rune and size in bytes
		r, size := utf8.DecodeRuneInString(word)
		fmt.Printf("%sFirst rune '%c' | Size: %d byte(s)\n", strings.Repeat(" ", 2), r, size)

		// print the last rune and size in bytes
		r, size = utf8.DecodeLastRuneInString(word)
		fmt.Printf("%sLast rune '%c' | Size: %d byte(s)\n", strings.Repeat(" ", 2), r, size)
	
		// slice and print the first two runes
		_, first := utf8.DecodeRuneInString(word)
		_, second := utf8.DecodeRuneInString(word[first:])
		fmt.Printf("%sFirst two runes: %s\n", strings.Repeat(" ", 2), word[:first+second])

		// slice and print the last two runes
		_, first = utf8.DecodeLastRuneInString(word)
		_, second = utf8.DecodeLastRuneInString(word[:len(word)-first])
		fmt.Printf("%sLast two runes: %s\n", strings.Repeat(" ", 2), word[len(word)-(first+second):])

		// convert string to rune and print the first and last two runes again
		rword := []rune(word)
		fmt.Printf("%sFirst two runes from rune slice: %c\n", strings.Repeat(" ", 2), rword[:2])
		fmt.Printf("%sLast two runes from rune slice: %c\n", strings.Repeat(" ", 2), rword[len(rword)-2:])		
	}

}