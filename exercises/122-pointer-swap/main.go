package main

import "fmt"

func main() {
	// Create a function that can swap 2 float variables
	// through their pointer addresses
	a, b := 1.23, 4.56
	fmt.Printf("Original address of 'a': %p 'b': %p\n", &a, &b)
	fmt.Printf("Before 'a'=%.2f 'b'=%.2f\n", a, b)
	valueSwapper(&a, &b)
	// confirm that both values have been swapped
	fmt.Printf("After 'a'=%.2f 'b'=%.2f\n", a, b)

	pa, pb := &a, &b
	fmt.Printf("Before address swap 'a'=%p,value %#v 'b'=%p value %#v\n", pa, *pa, pb, *pb)
	pa, pb = addressSwapper(pa, pb)
	fmt.Printf("After address swap 'a'=%p,value %#v 'b'=%p value %#v\n", pa, *pa, pb, *pb)
}

func valueSwapper(s, d *float64) {
	*s, *d = *d, *s
}

func addressSwapper(s, d *float64) (*float64, *float64) {
	return d, s
}
