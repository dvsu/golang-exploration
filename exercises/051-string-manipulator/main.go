package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args

	if len(args) != 3 {
		fmt.Println("[command] [string]\n\nAvailable commands: lower, upper and title")
		return
	}

	command, str := args[1], args[2]

	switch command {
	case "lower":
		fmt.Printf("%s\n", strings.ToLower(str))
	case "upper":
		fmt.Printf("%s\n", strings.ToUpper(str))
	case "title":
		fmt.Printf("%s\n", strings.Title(str))
	default:
		fmt.Printf("Unknown command: %q\n", command)
	}
}
