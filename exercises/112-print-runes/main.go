package main

import (
	"fmt"
	"strings"
)

func main() {
	// loop over 'word' and print its runes as decimals, hexadecimals and binary
	// Then, use the output values and print the word manually as runes, decimals and hexadecimals

	const word = "console"

	fmt.Printf("%-12s %-12s %-12s %-12s\n", "Rune", "Decimal", "Hexadecimal", "Binary")
	fmt.Printf("%s\n", strings.Repeat("=", 51))
	for _, r := range word {
		fmt.Printf("%-12c %-12[1]d %-12[1]X %-12[1]b\n", r)
	}

	// print manually using rune
	fmt.Printf("Runes: %s\n", string([]byte{'c', 'o', 'n', 's', 'o', 'l', 'e'}))

	// print manually using decimals
	fmt.Printf("Decimals: %s\n", string([]byte{99, 111, 110, 115, 111, 108, 101}))

	// print manually using hexadecimals
	fmt.Printf("Hexadecimals: %s\n", string([]byte{0x63, 0x6F, 0x6E, 0x73, 0x6F, 0x6C, 0x65}))
}