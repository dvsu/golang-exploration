package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// assume the os.Args[1] = Buenosdías
	msg := os.Args[1]
	fmt.Println(msg + strings.Repeat("*", utf8.RuneCountInString(msg)))

}
