// an improvement to exercise 064
//
// this time, we will overcome the difficulty of bigger
// guessed number by dynamically increasing the number of
// guesses. Therefore, bigger number has higher chance to win

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var (
		s       string
		n       int
		lim     int = 10
		maxTurn int = 10
		err     error
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

	if n > 10 {
		lim = n
	}

	// this part is the fix
	if n > 20 {
		maxTurn += n / 5
	}

	for i := 1; i <= maxTurn; i++ {
		pick := rand.Intn(lim) + 1

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
