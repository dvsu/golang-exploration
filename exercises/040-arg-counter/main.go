package main

import (
	"fmt"
	"os"
)

func pluralVerbCountNoun(l int) (string, string) {
	if l == 0 {
		return "is no", "s"
	} else if l == 1 {
		return "is one", ""
	} else {
		return "are more than one", "s"
	}
}

func main() {
	fmt.Println(len(os.Args))

	l := len(os.Args) - 1
	vbe, p := pluralVerbCountNoun(l)

	fmt.Printf("There %s argument%s\n", vbe, p)

}
