package main

import (
	"encoding/json"
	"fmt"
)

const data = `
[
  {
    "id": 1,
    "name": "need for speed",
    "genre": "racing",
    "price": 39.99
  },
  {
    "id": 2,
    "name": "simcity",
    "genre": "simulation",
    "price": 12.99
  },
  {
    "id": 3,
    "name": "minecraft",
    "genre": "sandbox",
    "price": 8.99
  }
]`

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

func load() (games []game) {
	var gamesjson []gamejson

	// use Unmarshal() to decode JSON
	// a function cannot change a variable directly
	// To allow function to change variable, we have to you pointer '&'
	// & -> address operator; it finds the memory address of a value/ variable
	if err := json.Unmarshal([]byte(data), &gamesjson); err != nil {
		fmt.Println("Unable to load JSON data:", err)
		return
	}

	for _, g := range gamesjson {
		games = append(games, game{
			item: item{
				id:    g.ID,
				name:  g.Name,
				price: g.Price,
			},
			genre: g.Genre,
		})
	}
	return
}

func newGame(id int, price float64, name, genre string) game {
	return game{
		item: item{
			id: id,
			name: name,
			price: price,
		},
		genre: genre,
	}
}

func addGame(gs []game, g game) []game {
	return append(gs, g)
}

func printGame(g game) {
	fmt.Printf("Name: %s\nID: %d\nGenre:%s\nPrice: %.2f\n", g.name, g.id, g.genre, g.price)
}

func saveGames(gs []game) error {
	gamesjson := make([]gamejson, 0, len(gs))
	for _, g := range gs {
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
		return err
	}
	fmt.Println(string(out))
	return nil
}
