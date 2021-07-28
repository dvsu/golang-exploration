package main

import "fmt"

func main() {
	firstName := "Adam"
	lastName := "Smith"

	fullName := "%s %s\n"
	fmt.Printf(fullName, firstName, lastName)
}
