package main

import (
	"fmt"
	"os"
	"strconv"
)

const usage = "Usage: go run main.go [value in USD]"

func main() {
	// currency exchange ratio using keyed elements
	const (
		EUR = iota
		GBP
		JPY
		AUD
	)

	forex := [...]float64{
		GBP: 0.71906,
		EUR: 0.84255,
		AUD: 1.3618,
		JPY: 109.68,
	}

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println(usage)
		return
	}

	if usd, err := strconv.ParseFloat(args[0], 64); err != nil {
		fmt.Println(usage)
		return
	} else {
		fmt.Printf("USD %.2f = GBP %.5f\n", usd, forex[GBP]*usd)
		fmt.Printf("USD %.2f = EUR %.5f\n", usd, forex[EUR]*usd)
		fmt.Printf("USD %.2f = AUD %.4f\n", usd, forex[AUD]*usd)
		fmt.Printf("USD %.2f = JPY %.2f\n", usd, forex[JPY]*usd)
	}
}
