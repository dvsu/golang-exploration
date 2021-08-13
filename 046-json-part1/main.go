package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// JSON -> JSON string = encode (marshals in Go)
	// JSON string -> JSON = decode (unmarshals in Go)

	/*
					                         key name
					                     (usually denotes  .------ value
										  package name)   /
					type person struct {        |        /
						                        V       V
						Name       string     `json:"fullname"`
						ID         int        `json:"-"`      <--------------- omit ID field during encoding
						Assets     assets     `json:"ast,omitempty"`
					}                                ^          ^
						 ^           ^              /      ^     \
					field name   field type        /   field tag  '-------- omit 'Assets' field if
					    ^                         / (string metadata        the value is empty (zero-value)
					   /	  change the    -----'    to a field)
					  /      encoded field
					 /	    name to 'ast'
					/		 instead of
				   / 		 'Assets'
				  /
		The first letter has to be capitalized
		to indicate the field is the exported field.
		It is important because json package
		only encodes exported fields
	*/

	type players map[int]string

	type team struct {
		Name     string  `json:"name"`
		Nickname string  `json:"-"`  // this field will be excluded during encoding
		Ground   string  `json:"ground"`
		Players  players `json:"players,omitempty"`
	}

	teams := []team{
		{
			Name:   "Liverpool",
			Ground: "Anfield",
			Players: players{
				1: "Alisson",
				3: "Fabinho",
				4: "Virgil van Dijk",
				5: "Ibrahima Konaté",
				6: "Thiago Alcântara",
			},
		},
		{
			Name:   "Manchester United",
			Ground: "Old Trafford",
			Players: players{
				1: "David de Gea",
				2: "Victor Lindelöf",
				3: "Eric Bailly",
				4: "Phil Jones",
				5: "Harry Maguire",
			},
		},
	}

	// out, err := json.Marshal(teams)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(out))

	// use MarchalIndent to prettyprint
	out, err := json.MarshalIndent(teams, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))	
}