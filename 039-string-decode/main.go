package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {

	text := `A vuestra espada no igualó la mía,		
Febo español, curioso cortesano,		
Ni a la alta gloria de valor mi mano,		
Que rayo fue do nace y muere el día.`
	
	fmt.Println("for-loop without decoding")
	// see what happen if the text is printed byte by byte
	for i := 0; i < len(text); i++ {
		fmt.Printf("%c", text[i])
	}
	fmt.Printf("\n\n")
	// the text is corrupted

	fmt.Println("Decode using for-loop + DecodeRuneInString()")
	// decode runes manually using for loop + DecodeRuneInString()
	for i := 0; i < len(text); {
		r, size := utf8.DecodeRuneInString(text[i:])
		fmt.Printf("%c", r)
		i += size
	}
	fmt.Printf("\n\n")

	fmt.Println("Decode using for-range loop")
	// for-range loop automatically decodes runes in string value
	for _, r := range text {
		fmt.Printf("%c", r)
	}
	fmt.Printf("\n\n")

	// convert string to byte slice, so the value can be overwritten
	word := []byte("ängsmarker") // 'meadows' in Swedish

	// extract the first UTF-8 encoding in p and return the rune and its width in bytes
	i, size := utf8.DecodeRune(word)

	fmt.Printf("%c %d\n", i, size) // i = 'ä', size = 2 bytes

	copy(word[:size], bytes.ToUpper(word[:size]))
	fmt.Printf("%s = % [1]x\n", word)
}