package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	id    int
	name  string
	price float64
}

type game struct {
	item
	genre string
}

type gamejson struct {
	ID	int      	`json:"id"`
	Name string		`json:"name"`
	Genre string  	`json:"genre"`
	Price float64 	`json:"price"`
}

func main() {
	// similar to exercise 117, but add 'save' command.
	// 'save' command will encode the 'games' to 
	// json and pretty print it

	games := []game{
		{
			item: item{
				id:    1,
				name:  "need for speed",
				price: 39.99,
			},
			genre: "racing",
		},
		{
			item: item{
				id:    2,
				name:  "simcity",
				price: 12.99,
			},
			genre: "simulation",
		},
		{
			item: item{
				id:    3,
				name:  "minecraft",
				price: 8.99,
			},
			genre: "sandbox",
		},
	}

	commands := []string{
		"list",
		"quit",
		"save",
		"getGameByID={ID_NUMBER}",
		"getGamesByGenre={GENRE}",
	}

	usage := "Available commands:\n"
	for _, c := range commands {
		usage += fmt.Sprintf("    %s\n", c)
	}

	usage += "\nExample:\n    getGameByID=2\n"

	var (
		id int
		err error
	)

queries:
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {

		input := in.Bytes()

		query := strings.Split(string(input), "=")

		switch query[0] {
		case "quit":
			return
		case "list":
			for _, g := range games {
				fmt.Println(g.name)
			}
		case "save":
			gamesjson := make([]gamejson, 0, len(games))
			for _, g := range games {
				gamesjson = append(gamesjson, gamejson{
					ID: g.id,
					Name: g.name,
					Genre: g.genre,
					Price: g.price,
				})
			}
			out, err := json.MarshalIndent(gamesjson, "", "  ")
			if err != nil {
				fmt.Println(err)
				return
			}
		
			fmt.Println(string(out))

		case "getGameByID":
			if len(query) != 2 {
				fmt.Println(usage)
				continue queries
			}

			if id, err = strconv.Atoi(query[1]); err != nil {
				fmt.Println(usage)
				continue queries
			}

			for _, g := range games {
				if id == g.id {
					fmt.Printf("Name: %s\nGenre: %s\nPrice: %.2f\n", g.name, g.genre, g.price)
					continue queries	
				}
			}
			fmt.Printf("There is no games with ID '%d'. Try again\n", id)

		case "getGamesByGenre":
			if len(query) != 2 {
				fmt.Println(usage)
				continue queries
			}

			for _, g := range games {
				if query[1] == g.genre {
					fmt.Printf("Name: %s\nID: %d\nPrice: %.2f\n", g.name, g.id, g.price)
					continue queries	
				}
			}
			fmt.Printf("There is no games with genre '%s'. Try again\n", query[1])
						
		default:
			fmt.Println(usage)
		}
	}
}
