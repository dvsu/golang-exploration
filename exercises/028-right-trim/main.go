package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// count number of characters, ignore white space. expected 6
	t := "	por qué         "
	fmt.Println(utf8.RuneCountInString(strings.TrimSpace(strings.ReplaceAll(t, " ", ""))))
}
