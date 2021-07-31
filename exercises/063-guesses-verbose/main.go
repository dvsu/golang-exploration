package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const maxTurn = 10

func main() {
	rand.Seed(time.Now().UnixNano())

	var (
		s   string
		n   int
		err error
	)

	args := os.Args[1:]

	if len(args) > 2 || len(args) == 0 {
		fmt.Println("Please guess a number")
		return
	}

	opt := len(args) == 2

	if opt && args[0] != "-v" {
		fmt.Println("Option is unknown. Please try -v to generate verbose output")
		return
	}

	switch opt {
	case true:
		s = args[1]
	case false:
		s = args[0]
	}

	if n, err = strconv.Atoi(s); err != nil {
		fmt.Printf("Input is not a number. Try again!")
		return
	}

	for i := 1; i <= maxTurn; i++ {
		pick := rand.Intn(n) + 1

		if opt {
			fmt.Printf("%d ", pick)
		}
		if pick != n {
			continue
		}

		if i == 1 {
			fmt.Println("Achievement Unlocked! First-turn Winner!")
		} else {
			fmt.Println("You win!")
		}
		return
	}
	fmt.Println("You lose! Try again!")
}
