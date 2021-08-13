package main

import "fmt"

type item struct {
	id    int
	name  string
	price float64
}

type game struct {
	item
	genre string
}

func main() {
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

	fmt.Printf("%#v\n", games)
}