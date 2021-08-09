package main

import "fmt"

func main() {
	// convert 'words' to byte slice and append to 'bwords'
	// print 'bwords' as byte and string slices

	words := []string{
		"Lorem ipsum",
		"dolor sit amet",
		"consectetur adipiscing elit",
		"Sed sagittis vestibulum urna non congue",
	}

	var bwords [][]byte

	for i, word := range words {
		bwords = append(bwords, []byte(word))
		fmt.Printf("%d\n", bwords[i])
	}

	for _, b := range bwords {
		fmt.Printf("%s\n", b)
	}
}