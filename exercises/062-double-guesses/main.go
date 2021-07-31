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
		n    int
		m    int
		lim  int
		nErr error
		mErr error
	)

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Please guess 2 numbers")
		return
	}

	if n, nErr = strconv.Atoi(args[0]); nErr != nil {
		fmt.Printf("First input is not a number. Try again!")
		return
	}

	lim = n

	if m, mErr = strconv.Atoi(args[1]); mErr != nil {
		fmt.Printf("Second input is not a number. Try again!")
		return
	}

	// handle negative number
	if m < 1 || n < 1 {
		fmt.Println("Please guess 2 positive numbers")
		return
	}

	if m > n {
		lim = m
	}

	for i := 1; i <= maxTurn; i++ {
		pick := rand.Intn(lim) + 1

		if pick != n && pick != m {
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
