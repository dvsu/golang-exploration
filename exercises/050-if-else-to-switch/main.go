package main

import (
	"fmt"
	"os"
)

const (
	usage       = "Usage: [username] [password]"
	errUser     = "Access denied for %q.\n"
	errPwd      = "Invalid password for %q.\n"
	accessOK    = "Access granted to %q.\n"
	user, user2 = "alice", "bob"
	pass, pass2 = "2468", "1357"
)

func main() {

	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	// Refactor the following codes to switch statement

	// if u != user && u != user2 {
	// 	fmt.Printf(errUser, u)
	// } else if u == user && p == pass {
	// 	fmt.Printf(accessOK, u)
	// } else if u == user2 && p == pass2 {
	// 	fmt.Printf(accessOK, u)
	// } else {
	// 	fmt.Printf(errPwd, u)
	// }

	switch {
	case u != user && u != user2:
		fmt.Printf(errUser, u)
	case u == user && p == pass:
		fallthrough
	case u == user2 && p == pass2:
		fmt.Printf(accessOK, u)
	default:
		fmt.Printf(errPwd, u)
	}
}
