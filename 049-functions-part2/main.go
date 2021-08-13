package main

import "fmt"

type character struct {
	hp, mp, level int
	name          string
}

func main() {
	// Go is a pass-by-value language
	// It means everything that is passed to
	// a function is not the original data,
	// but the copied version.

	// Therefore, it is necessary to return the modified data
	// from a function and reassign the value to make sure
	// the original data is equal to the updated data.
	var characters []character

	// this one will not be included in the 'characters' slice because
	// the returned value is not assigned to it
	createCharacter("Alpha")

	characters = append(characters, createCharacter("Bravo"), createCharacter("Charlie"))

	fmt.Printf("%#v\n", characters)

	// However, the following function will update the original slice content
	// because all points to the same memory location.
	// Changing the value of the copied version will also change
	// the original slice, even the function does not return anything.
	levelUpAllCharacters(characters)

	fmt.Println("After level up")
	fmt.Printf("%#v\n", characters)
}

func createCharacter(name string) (c character) {
	c = character{
		hp:    300,
		mp:    20,
		level: 1,
		name:  name,
	}
	return
}

func levelUpAllCharacters(ch []character) {
	for i := range ch {
		ch[i].level++
	}
}
