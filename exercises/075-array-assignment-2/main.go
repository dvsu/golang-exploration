package main

import (
	"fmt"
	"strings"
)

func main() {
	// assign books array to the other 2 arrays, but
	// change the letter case to upper and lower case
	// respectively

	books := [...]string{
		"Chronicles of Avonlea",
		"Rainbow Valley",
		"Anne's House of Dreams",
		"Rilla of Ingleside",
		"The Story of My Life",
	}

	var (
		bUpper [5]string
		bLower [5]string
	)

	for i, e := range books {
		bUpper[i], bLower[i] = strings.ToUpper(e), strings.ToLower(e)
	}

	fmt.Printf("%q\n", books)
	fmt.Printf("%q\n", bUpper)
	fmt.Printf("%q\n", bLower)
}
