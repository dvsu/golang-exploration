package main

import (
	"fmt"
	"time"
)

func main() {
	sysOs := "Windows 10"

	switch sysOs {
	case "Windows 7", "Windows 10":
		fmt.Println("Windows")
	case "MacOS":
		fmt.Println("MacOS")
	case "Ubuntu", "CentOS", "Manjaro":
		fmt.Println("Linux")
	default:
		fmt.Println("Unknown OS")
	}

	// boolean expression
	// the default value of switch condition is true
	// instead of typing 'switch true {...}'
	// try 'switch {...}' for more concise code

	t := 26.7

	switch {
	case t > 25.0:
		fmt.Println("Hot")
	case t < 16.0:
		fmt.Println("Cold")
	default:
		fmt.Println("Comfort")
	}

	// fallthrough statement can be used to ignore the next case statement and execute the action

	cm := 106

	switch {
	case cm > 115:
		fmt.Print("X")
		fallthrough
	case cm > 110:
		fmt.Print("X")
		fallthrough
	case cm > 105:
		fmt.Print("X")
		fallthrough
	case cm > 100:
		fmt.Print("L\n")
	}

	// short switch using simple statement
	switch h := 53.2; {
	case h > 65.0:
		fmt.Println("Too humid")
	case h < 45.0:
		fmt.Println("Too dry")
	default:
		fmt.Println("Perfect humidity")
	}

	// parts of a day using switch statement
	switch t := time.Now().Hour(); {
	case t > 17:
		fmt.Println("Good evening")
	case t > 11:
		fmt.Println("Good afternoon")
	case t > 5:
		fmt.Println("Good morning")
	default:
		fmt.Println("Good night")
	}
}
