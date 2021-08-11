package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

const houses = "Available house names: gryffindor, hufflepuf, ravenclaw, slytherin"

func main() {
	// get a query word from command line and print the result
	// in ascending order
	data := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
	}

	// delete 'bobo'
	delete(data, "bobo")

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf("Usage: go run main.go {house_name}\n%s\n", houses)
		return
	}

	if v, e := data[args[0]]; !e {
		fmt.Printf("House name %q does not exist\n%s\n", args[0], houses)
		return
	} else {
		sorted := append([]string(nil), v...)
		sort.Strings(sorted)
		fmt.Printf("===== %s students =====\n", args[0])
		for _, s := range sorted {
			fmt.Printf("%s+ %s\n", strings.Repeat(" ", 4), s)
		}
	}
}