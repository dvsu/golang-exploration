package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	usage = `
Available commands:
    list
    quit
    save
    add {ID|int} {NAME|string} {GENRE|string} {PRICE|float64}
    getGameByID {ID|int}
    getGamesByGenre {GENRE|string}

Examples:
    getGameByID 2
    getGamesByGenre racing
`
)

func runCmd(input string, games []game) (content []game, cont bool) {

	content, cont = nil, true

	query := strings.Fields(input)

	if len(query) == 0 {
		return
	}

	switch query[0] {
	case "quit":
		cont = cmdQuit()
		return
	case "list":
		cmdList(games)
	case "save":
		err := saveGames(games)
		if err != nil {
			fmt.Println(err)
		}
	case "add":
		if len(query) != 5 {
			fmt.Println("Information is missing")
			fmt.Println(usage)
			return
		}

		id, err := strconv.Atoi(query[1])
		if err != nil {
			fmt.Println("Wrong ID format")
			fmt.Println(usage)
			return
		}
		
		name, genre := query[2], query[3]

		price, err := strconv.ParseFloat(query[4], 64)
		if err != nil {
			fmt.Println("Wrong price format")
			fmt.Println(usage)
			return
		}

		return addGame(games, newGame(id, price, name, genre)), true

	case "getGameByID":
		if len(query) != 2 {
			fmt.Println(usage)
			return
		}

		id, err := strconv.Atoi(query[1])
		if err != nil {
			fmt.Println("Wrong ID format")
			fmt.Println(usage)
			return
		}

		for _, g := range games {
			if id == g.id {
				fmt.Printf("Name: %s\nGenre: %s\nPrice: %.2f\n", g.name, g.genre, g.price)
				return	
			}
		}
		fmt.Printf("There is no games with ID '%d'. Try again\n", id)

	case "getGamesByGenre":
		if len(query) != 2 {
			fmt.Println(usage)
			return
		}

		count := 0
		for _, g := range games {
			if query[1] == g.genre {
				fmt.Printf("Name: %s\nID: %d\nPrice: %.2f\n", g.name, g.id, g.price)
				count++
			}
		}
		if count == 0 {
			fmt.Printf("There is no games with genre '%s'. Try again\n", query[1])
		}
	default:
		fmt.Println(usage)
	}
	return
}

func cmdQuit() bool {
	fmt.Println("Exiting the program...")
	return false
}

func cmdList(games []game) {
	for _, g := range games {
		printGame(g)		
	}
}
