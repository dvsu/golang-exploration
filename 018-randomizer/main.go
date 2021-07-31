package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	target := 20

	for n := 0; n != target; {
		n = rand.Intn(target + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
