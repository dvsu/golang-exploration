package main

import (
	"bytes"
	"fmt"
)

func main() {
	// append png to header, compare and check whether both slices are equal

	png, header := []byte{'P', 'N', 'G'}, []byte{}

	header = append(header, png...)

	if bytes.Equal(png, header) {
		fmt.Println("Equal")
	}

}
