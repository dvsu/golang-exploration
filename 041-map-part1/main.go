package main

import (
	"fmt"
	"strings"
)

func main() {
	// map can be read even without initializing it
	// map syntax
	//   map[key-data-type]value-data-type
	var dict map[string]string

	fmt.Printf("Map %#v Key count: %d\n", dict, len(dict)) // 0

	// map will return zero value if the key does not exist in the map.
	// the zero value deponds on the data type of the value
	// e.g. 0 for int
	//      "" for string
	//      false for bool
	fmt.Printf("Map 'dict' with key %q has a value equal to %q\n", "something", dict["something"])

	// map may also have slice as value
	// 'collection' map has 'string' key and 'float64' slice as value
	var collection map[string][]float64
	fmt.Printf("Map %#v Key count: %d\n", collection, len(collection)) // 0

	// map key should be a comparable type, such as int, bool, string, etc.
	// slice, map, and array are not comparable, so they cannot be used as key
	// it is not recommended to use float as key type because it may lead to inaccuracy

	// a map has to be initialized first before it can be used, e.g. adding key-value pairs
	// we can use short declaration to initialize a map
	// 'dictEnEs' is initialized with a few key-value pairs
	dictEnEs := map[string]string{
		"please": "por favor",
		"thank you": "gracias",
		"sorry": "la siento",
		"yes": "sí",
		"no": "no",
	}
	fmt.Printf("Map %#v Key count: %d\n", dictEnEs, len(dictEnEs)) // 5

	// add key-value pairs
	dictEnEs["hello"] = "hola"
	dictEnEs["goodbye"] = "adiós"
	fmt.Printf("Map %#v after adding a few key-value pairs. Key count: %d\n", dictEnEs, len(dictEnEs)) // value of map is not nil

	// to check whether a key exists in a map
	val, exist := dictEnEs["run"]; // val = '', exist = false
	fmt.Printf("Key 'run' exists in 'dictEnEs' map: %t. Value: %q\n", exist, val) // false

	q := "please"
	// alternatively, we can also use short-if
	if v, e := dictEnEs[q]; e {
		fmt.Printf("%q means %q in Spanish\n", q, v)
	}

	// We can also use for-range loop to loop over a map
	// the order of printed result does change sometimes because
	// map is unindexed and more suitable for fast-lookup application
	fmt.Printf("%-12q %-12q\n", "English", "Spanish")
	fmt.Printf("%s\n", strings.Repeat("=", 25))
	for k, v := range dictEnEs {
		fmt.Printf("%-12q %-12q\n", k, v)
	}

	// However, the print output is always sorted by key
	// So the string output can be used to perform equality check
	original := fmt.Sprintf("%s", dictEnEs)

	// let's shuffle the order
	dictEnEs2 := map[string]string{
		"sorry": "la siento",
		"yes": "sí",
		"thank you": "gracias",
		"no": "no",
		"hello": "hola",
		"goodbye": "adiós",
		"please": "por favor",
	}

	copied := fmt.Sprintf("%s", dictEnEs2)

	if original == copied {
		fmt.Println("Both maps are equal")
	}

	// proof
	fmt.Printf("'original' %s\n", dictEnEs)
	fmt.Printf("'shuffled' %s\n", dictEnEs2)	
}