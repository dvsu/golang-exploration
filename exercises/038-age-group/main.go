package main

import "fmt"

func main() {
	// age category using if-else statement
	age := 3

	if age > 45 {
		fmt.Println("Old adults")
	} else if age > 30 {
		fmt.Println("Middle-aged adults")
	} else if age > 16 {
		fmt.Println("Young adults")
	} else if age > 2 {
		fmt.Println("Children")
	} else {
		fmt.Println("Babies")
	}
}
