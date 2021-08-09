package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// string values are read-only byte slices

	greeting := "buenos días ☀☀☀"

	fmt.Println(greeting[0]) // output is 98, numeric representation of 'b'
	fmt.Println(greeting[8]) // output is 195, numeric representation of 'í'

	// Since it is a read-only slice, we cannot modify the element
	// greeting[6] = '_' // this is invalid

	// Modifying an element is possible by converting the read-only byte slice to byte slice first.
	greetingBytes := []byte(greeting)
	
	greetingBytes[6] = '_' // this is valid

	// the original slice is untouched because the conversion creates a new backing array
	// and copies the bytes of string to new one
	fmt.Printf("Original slice 'greeting': %s\n", greeting) // "buenos días"
	fmt.Printf("New slice 'greetingBytes': %s\n", greetingBytes) // "buenos_días"

	// Similarly, conversion from byte to string also creates a new backing array
	greetingString := string(greetingBytes)
	fmt.Printf("Original byte slice 'greetingBytes': %s\n", greetingBytes) // "buenos_días"
	fmt.Printf("After conversion byte -> string: %s\n", greetingString) // "buenos_días"

	// it also means the slice after conversion from byte to string becomes read-only
	// this is valid
	greetingBytes[0] = 'J'
	// but this is invalid
	// greetingString[0] = 'J'

	// when looping over a string, the for loop will loop over the rune
	// instead of byte-per-byte. Therefore, looping over 'greeting' variable
	// will generate the following output
	// [0 ] 62   = 'b' | as single byte rune: b
	// [1 ] 75   = 'u' | as single byte rune: u
	// [2 ] 65   = 'e' | as single byte rune: e
	// [3 ] 6E   = 'n' | as single byte rune: n
	// [4 ] 6F   = 'o' | as single byte rune: o
	// [5 ] 73   = 's' | as single byte rune: s
	// [6 ] 20   = ' ' | as single byte rune:
	// [7 ] 64   = 'd' | as single byte rune: d
	// [8 ] C3AD = 'í' | as single byte rune: Ã <-- 2 bytes for non-English characters.
	// [10] 61   = 'a' | as single byte rune: a    see how the single byte rune outputs the wrong character
	// [11] 73   = 's' | as single byte rune: s    because the correct way to extract the rune is greeting[8:10]
	// [12] 20   = ' ' | as single byte rune:
	// [13] E29880 = '☀' | as single byte rune: â <-- this is wrong because the emoji should be 3 bytes long
	// [16] E29880 = '☀' | as single byte rune: â <-- this
	// [19] E29880 = '☀' | as single byte rune: â <-- and this as well

	for i, r := range greeting {
		fmt.Printf("[%-2d] %-4X = %q | as single byte rune: %c\n", i, string(r), r, greeting[i])
	}

	fmt.Printf("This is wrong: %c\n", greeting[8])
	fmt.Printf("This is correct: %s\n", greeting[8:10])

	fmt.Printf("This is wrong: %c\n", greeting[13])
	fmt.Printf("This is correct: %s\n", greeting[13:17])

	// there is actually a way to convert the string to rune slice, but the operation is costly.
	greetingRune := []rune(greeting)

	fmt.Printf("%c\n", greetingRune)
	// 15 characters, even though the actual byte size is larger
	fmt.Printf("Rune count/length: %d characters\n", len(greetingRune))
	// it also fixes the previous problem, where we have to slice 3 bytes to extract an emoji
	// here we can do it right away
	fmt.Printf("%c\n", greetingRune[14])  // print the emoji

	// as mentioned, the tradeoff is the size of rune slice, which allocates 4 bytes per element
	// let's compare the size
	fmt.Printf("Size of the original string: %d bytes\n", int(unsafe.Sizeof(greeting)))
	fmt.Printf("Size of rune slice: %d bytes\n", int(unsafe.Sizeof(greetingRune[0]))*len(greetingRune))
}