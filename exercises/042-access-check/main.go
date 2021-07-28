package main

import (
	"fmt"
	"os"
)

func main() {
	const (
		user = "anonymous"
		pass = "superS3cr3T"
	)
	l := len(os.Args) - 1
	if l != 2 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	u, p := os.Args[1], os.Args[2]

	if u == user && p == pass {
		fmt.Printf("Access granted to %q\n", u)
	} else if u == user && p != pass {
		fmt.Printf("Invalid password for %q\n", u)
	} else {
		fmt.Printf("Access denied for %q\n", u)
	}

}
