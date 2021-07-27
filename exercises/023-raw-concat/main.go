package main

import (
	"fmt"
	"os"
)

func main() {

	msg := `hi ` + os.Args[1] + `!
How are you?`

	fmt.Println(msg)
}
