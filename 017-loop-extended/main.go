package main

import (
	"fmt"
	"strings"
)

func main() {
	t := `Half-past six on a July morning! 
	The Mariposa Belle is at the wharf, decked in flags, 
	with steam up ready to start.`

	replacer := strings.NewReplacer("!", "", ",", "", ".", "")

	text := strings.Fields(replacer.Replace(t))

	for i := 0; i < len(text); i++ {
		fmt.Println(text[i])
	}

	// the for loop can be simplified using 'for range'
	fmt.Println("Using 'for range' loop to print both index and value")
	for i, v := range text {
		fmt.Println(i, v)
	}

	// if the value variable isn't needed, blank identifier can be used to replace the original variable
	fmt.Println("'for range' loop with blank identifier. It will print the index and ignore the value")
	for i, _ := range text {
		fmt.Println(i)
	}

	// even simpler, the blank identifier can be totally omitted if it isn't needed
	fmt.Println("Simple 'for range' loop. It will only print the index, same as the previous loop")
	for i := range text {
		fmt.Println(i)
	}

	l := []string{"omit_me", "next_one", "still_continue", "the_end"}

	for i, s := range l[1:] {
		fmt.Printf("%d: %q\n", i, s)
	}
}
