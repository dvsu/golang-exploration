package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	// assume the os.Args[1] = Buenosdías
	// the total character count should be 10
	fmt.Println(utf8.RuneCountInString(os.Args[1]))
}
