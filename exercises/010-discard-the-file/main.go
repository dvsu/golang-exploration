package main

import (
	"fmt"
	"path"
)

func main() {
	dir, _ := path.Split("secret/file.text")
	fmt.Println(dir)
}
