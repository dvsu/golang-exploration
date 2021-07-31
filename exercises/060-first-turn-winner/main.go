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
		n   int
		err error
	)

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please guess a number")
		return
	}

	if n, err = strconv.Atoi(args[0]); err != nil {
		fmt.Printf("Input is not a number. Try again!")
		return
	}

	for i := 1; i <= maxTurn; i++ {
		pick := rand.Intn(n + 1)
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
