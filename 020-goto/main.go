package main

import "fmt"

func main() {
	i := 5

	// Similar to labeled break and continue, goto can be used to
	// branch to particular location inside the function body
	// though the functionality is more restricted, e.g.
	// placing goto label before variable declaration is prohibited

	// simple goto to mimick loop (not recommended, example only)
loop:
	if i < 10 {
		fmt.Println(i)
		i++
		goto loop
	}
}
