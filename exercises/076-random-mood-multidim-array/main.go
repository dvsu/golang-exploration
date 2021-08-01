package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const usage = "Usage: go run main.go [your name] [positive|negative]"

func main() {

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	name, mood := args[0], args[1]

	moods := [...][3]string{
		{"cheerful", "humorous", "thankful"},
		{"gloomy", "pessimistic", "sad"},
	}

	rand.Seed(time.Now().UnixNano())
	pos := rand.Intn(len(moods[0]))
	var cat int

	switch mood {
	case "positive":
		cat = 0
	case "negative":
		cat = 1
	default:
		fmt.Println(usage)
		return
	}

	fmt.Printf("%s feels %s\n", name, moods[cat][pos])
}
