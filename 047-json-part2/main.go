package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type players map[int]string

type team struct {
	Name     string
	Nickname string
	Ground   string
	Players  players
}

func main() {
	// use the following command to feed JSON data
	// go run main.go < teams.json

	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	var teams []team

	// use Unmarshal() to decode JSON
	// a function cannot change a variable directly
	// To allow function to change variable, we have to you pointer '&'
	// & -> address operator; it finds the memory address of a value/ variable
	if err := json.Unmarshal(input, &teams); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", teams)
	fmt.Printf("%#v\n", teams[1])
}
