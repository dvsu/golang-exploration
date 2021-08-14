package main

import "fmt"

// method in Go
/*
   receiver name,
   usually abbreviated
   e.g. d for data
        c for count
      |     .-------------- receiver type
      |     |          .------------ method name
      V    V          V
func (r receiver) methodNmae() {
	// do something
}
*/

func main() {
	alpha := character{
		name: "Alpha",
		level: 1,
		hp: 300,
		mp:20,
	}

	bravo := character{
		name: "Bravo",
		level: 9,
		hp: 721,
		mp:35,
	}

	alpha.status()
	bravo.status()

	// we can also achieve the same result using method expression
	// Based on the behavior of the method, it is basically
	// a function, which takes receiver as its first argument
	// fmt.Println("Using method expression")
	// character.status(alpha)
	// character.status(bravo)

	// the difference: 
	// alpha.status() passes the receiver automatically
	// character.status(alpha) passes the receiver manually

	// method vs. function difference:
	//    method belongs to a type, function doesn't
	//    That's why calling status() without 'character' type
	//    will never work.
	//    function belongs to a package. It can be accessed
	//    package-wide.
	
	// The receiver is just a copy of the original values.
	// In order to modify the original value, the receiver
	// type has to be a pointer, i.e. points to the memory
	// location of original values.
	// The example syntax:
	// (&alpha).levelup()

	// After calling this method, Alpha will level up to level 2
	(&alpha).levelup()

	// Alternatively, we can also use this syntax.
	// Go will automatically insert the '&' operator
	// and take the address automatically 
	// when it sees pointer as the receiver type.
	// After calling this method, Alpha will level up to level 3
	alpha.levelup()

	fmt.Println("After calling levelup() method twice")
	alpha.status()

	// When one of the methods has pointer type as the receiver type
	// it would be good to change all method's receiver type to pointer
	// type.
}