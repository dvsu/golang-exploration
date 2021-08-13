package main

import "fmt"

func main() {
	// every pointer has its own type
	var n int = 128
	P := &n // the pointer type of P is *int
	// because it stores the memory
	// address of variable n that has
	// type int



	// we can also access the underlying value stored at address P
	// using *P (known as indirecting or dereferencing)
	V := *P // the value is stored in variable V
	// V will have different memory location


	fmt.Printf("Address of n: %p\n", P)
	fmt.Printf("Address of V: %p\n", &V)

	// because a pointer is also a value, it also has its memory location
	fmt.Printf("Address of P: %p\n", &P)

	fmt.Printf("Value of n: %d\n", n)
	fmt.Printf("Value of V: %d\n", V)	

	// we can also update the value indirectly using pointer
	// let's update the value of n to 256 using pointer
	*P = 256

	// check the result
	fmt.Printf("Value of n after update: %d\n", n)	

	// When a pointer is passed to a function
	// the it still points to the same memory location
	// but the value is stored at different memory location.
	// Every time a pointer is passed to a function.
	// the value will be created and released once
	// the function call is completed.

	// Because it still points to the same memory location,
	// we can indirectly modify the value at that address,
	// which usually cannot be done by normal variables
	modifier(P)
	fmt.Println("Now checking the value of n from main function...")

	// the address is still the same, but the value has been modified by
	// function modifier()
	fmt.Printf("Address of n: %p. Value of n now: %d\n", &n, n)
}

func modifier(i *int) {
	fmt.Printf("Address of i: %p. Value of i: %p\n", &i, i)
	fmt.Println("Now changing the value of n from a function...")
	*i = 512
}