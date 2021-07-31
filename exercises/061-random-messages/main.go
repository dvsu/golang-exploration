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

		switch rand.Intn(3) {
		case 0:
			fmt.Println("You win!")
		case 1:
			fmt.Println("Good job!")
		case 2:
			fmt.Println("You nailed it!")
		}
		return
	}

	switch rand.Intn(3) {
	case 0:
		fmt.Println("You lose!")
	case 1:
		fmt.Println("Almost there. Try again!")
	case 2:
		fmt.Println("Try again. Don't give up!")
	}
}
