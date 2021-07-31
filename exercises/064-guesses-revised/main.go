// a fix to variable guessing difficulty
//
// previously, a user who guesses a small number, e.g. 5
// may have easier difficulty because the random number
// is generated between 1 and that number
// To set a base difficulty, random number will be drawn
// between 1 and 10 for any guessed number given less than 10

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
		lim int = 10
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

	if n > 10 {
		lim = n
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
