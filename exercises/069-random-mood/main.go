package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// print random mood
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Usage: go run main.go [your name]")
		return
	}

	moods := [...]string{
		"cheerful",
		"reflective",
		"gloomy",
		"humorous",
		"thankful",
	}

	rand.Seed(time.Now().UnixNano())

	fmt.Printf("%s feels %s\n", args[0], moods[rand.Intn(len(moods))])
}
