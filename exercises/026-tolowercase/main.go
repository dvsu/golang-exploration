package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// convert the os.Args[1] value to lower case
	fmt.Println(strings.ToLower(os.Args[1]))
}
