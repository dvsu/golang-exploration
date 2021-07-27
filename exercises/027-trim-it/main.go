package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := `
	
	On New Year’s Eve we were all together in Uncle Alec’s kitchen,
which was tacitly given over to our revels during the winter evenings.
	`
	// remove all leading and trailing white space
	fmt.Println(strings.TrimSpace(msg))
}
