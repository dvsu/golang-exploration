package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	switch {
	case n == 2, n == 3:
		return true
	case n%2 == 0, n%3 == 0:
		return false
	}

	for i, w := 5, 2; i*i <= n; {
		if n%i == 0 {
			return false
		}
		i += w
		w = 6 - w
	}
	return true
}

func main() {

	var (
		i   int
		err error
	)

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Please give some numbers")
		return
	}

	for _, n := range args {
		if i, err = strconv.Atoi(n); err != nil {
			continue
		}
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
