package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	for {
		t := strings.Fields("\\ | / -")

		for _, s := range t {
			// fmt.Printf("\033[0;0H")
			fmt.Printf("\r%s Please wait. Processing...\n", s)
			time.Sleep(250 * time.Millisecond)
		}

	}
}
