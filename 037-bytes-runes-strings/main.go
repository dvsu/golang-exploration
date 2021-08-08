package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// rune literals are typeless
	// so they can be assigned to a variable with any numeric type

	var num1, num2 int

	// num1 = 85, num2 = 83
	num1, num2 = 'U', 'S'

	fmt.Println(num1, num2)

	// we can use %c verb to print runes
	// "U S"
	fmt.Printf("%c %c\n", num1, num2)

	// because it is numeric type, we can use 'for' loop to perform iteration
	// start = 65, stop = 122
	start, stop := 'A', 'z'

	// we will print it as a rune literal, decimal, hex, binary, and encoded
	fmt.Printf("%-12s %-12s %-12s %-12s %-12s\n", "Literal", "Decimal", "Hexadecimal", "Binary", "Encoded")
	fmt.Printf("%s\n", strings.Repeat("=", 60))
	for c := start; c <= stop; c++ {
		fmt.Printf("%-12c %-12[1]d %-12[1]X %-12[1]b % -12X \n", c, string(c))
	}

	// ASCII rune size is 1 byte (range 0 - 127)
	// ASCII is part of the Unicode Standard
	// 0 - 32 are control codes and they cannot be printed

	// let's see the representation if we extend the number beyond 127
	fmt.Printf("%-12s %-12s %-12s %-12s %-12s\n", "Literal", "Decimal", "Hexadecimal", "Binary", "Encoded")
	fmt.Printf("%s\n", strings.Repeat("=", 60))
	for c := 161; c <= 255; c++ {
		fmt.Printf("%-12c %-12[1]d %-12[1]X %-12[1]b % -12X \n", c, string(c))
	}

	// A byte (uint8) can store a number between 0 - 255
	// Code points that are in this range can fit into one byte
	// Literal à can fit into a single byte, as 224 is within
	// range of a single byte (0 - 255).
	// However, when it is UTF-8 encoded it cannot fit into
	// a single byte and requires 2 bytes (C3 A0)

	// remember this?
	fmt.Printf("à UTF-8 encoded: %d bytes\n", len("à"))
	fmt.Printf("à Count per string: %d byte\n", utf8.RuneCountInString("à"))

	// special characters require 3 bytes (UTF-8 encoded)
	fmt.Printf("%-18s %-18s %-18s %-18s %-128s\n", "Literal", "Decimal", "Hexadecimal", "Binary", "Encoded")
	fmt.Printf("%s\n", strings.Repeat("=", 60))
	for c := 9984; c <= 10175; c++ {
		fmt.Printf("%-18c %-18[1]d %-18[1]X %-18[1]b % -18X \n", c, string(c))
	}

	// let's compare again
	fmt.Printf("✅ UTF-8 encoded: %d bytes\n", len("✅"))
	fmt.Printf("✅ Count per string: %d byte\n", utf8.RuneCountInString("✅"))

	// similarly emojis require 4 bytes (UTF-8 encoded)
}
