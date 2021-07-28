package main

import "fmt"

func main() {
	value, alive := 3, true

	if value > 2 && alive {
		fmt.Println("You win!")
	} else if value == 0 && alive {
		fmt.Println("Invalid. Try again")
	} else {
		fmt.Println("You lose!")
	}
}
