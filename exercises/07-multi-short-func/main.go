package main

import "fmt"

func main() {
	_, n := multi()
	fmt.Println(n)
}

func multi() (int, int) {
	return 5, 4
}
