package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// expected "myprogram"
	// build using "go build -o myprogram.exe" on Windows
	_, file := filepath.Split(os.Args[0])
	fmt.Println(file)
}
