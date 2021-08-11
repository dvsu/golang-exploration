package main

import "fmt"

func main() {
	// Similar to slice, map does not actually store the key-value pairs.
	// Instead, they are actually stored in memory and the map header
	// contains the information that points to the memory location.
	// Therefore, copying a map then updating the value may also
	// change the value in original map.
	// Interestingly, the map value itself only contains and points to
	// the memory address of a map header.
	// map variable ---> map header ---> actual value in memory
	source := map[string]string{
		"sorry":     "la siento",
		"yes":       "sí",
		"thank you": "gracias",
		"no":        "no",
	}

	copied := source
	copied["please"] = "por favor"

	fmt.Printf("%#v\n", source) // now source also has "please": "por favor" in its map
								// because they point to the same memory location.
	fmt.Printf("%#v\n", copied) // Copying a map is equal to copying its memory location

	// there are 2 ways to create a new map
	// 1. using make() (can also preallocate length, but the length itself may not be directly allocated
	//					and will be decided by compiler and during runtime). 
	// 2. using short declaration and assign empty value
	// to create a new map, we can use make()
	another := make(map[string]string, len(source))

	// the length is 0 even though we give a fixed length value during declaration
	fmt.Printf("Length of 'another' map: %d\n", len(another))

	// store the reversed translation in 'another' map using for loop
	for k, v := range source {
		another[v] = k
	}

	// check the result
	fmt.Printf("'source' map content: %s\n", source)
	fmt.Printf("'another' map content: %s\n", another)

	// to delete a key-value pair, we can use delete()
	// It is also safe to call delete() even the key does not exist.
	// So, no preliminary checking is required.
	delete(source, "thank you")

	//final check
	fmt.Printf("'source' map content after removing 'thank you' key: %s\n", source)
	fmt.Printf("'another' map content is not affected by the removal: %s\n", another)

	// Assigning a nil value to a map will not destroy the map
	// The assignment only removes the pointer to the memory location
	// and the actual values still live in the memory.

	// To completely remove a map, use mapclear()
}
