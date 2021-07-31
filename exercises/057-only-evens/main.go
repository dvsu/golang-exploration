package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	var t int

	if len(args) != 3 {
		fmt.Println("Please give 2 numbers that represent minimum and maximum range")
		return
	}

	if min, minErr := strconv.Atoi(args[1]); minErr != nil {
		fmt.Println("Minimum number is not valid")
		return
	} else if max, maxErr := strconv.Atoi(args[2]); maxErr != nil {
		fmt.Println("Maximum number is not valid")
		return
	} else if min > max {
		fmt.Println("Minimum number is larger than maximum number")
		return
	} else {
		for i := min; i <= max; i++ {
			if i%2 != 0 {
				continue
			}
			t += i
			fmt.Printf("%d ", i)
			if i != max {
				fmt.Printf("+ ")
			}
		}
		fmt.Printf("= %d\n", t)

	}

}
