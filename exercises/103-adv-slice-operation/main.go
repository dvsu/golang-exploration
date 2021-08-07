package main

import "fmt"

func main() {
	// create a string slice with length and capacity of 5
	presidents := []string{"", "", "", "", ""}

	fmt.Printf("Length: %d Capacity: %d b.array %p 'presidents' %q\n", len(presidents), cap(presidents), presidents, presidents)

	// append 3 elements to it
	for _, s := range []string{"George Washington", "John Adams", "Thomas Jefferson"} {
		presidents = append(presidents, s)
	}

	fmt.Printf("Length: %d Capacity: %d b.array %p 'presidents' %q\n", len(presidents), cap(presidents), presidents, presidents)
	// the slice, however, has unexpected result ["", "", "", "", "", s, s, s]
	// expected [s, s, s, "", ""]

	// recreate new string slice with capacity of 5 using make() and fix the previous problem
	otherPresidents := make([]string, 0, 5)

	for _, s := range []string{"George Washington", "John Adams", "Thomas Jefferson"} {
		otherPresidents = append(otherPresidents, s)
	}

	// then reveal the whole length
	otherPresidents = otherPresidents[:5]

	fmt.Printf("Length: %d Capacity: %d b.array %p 'otherPresidents' %q\n", len(otherPresidents), cap(otherPresidents), otherPresidents, otherPresidents)

	// copy the first two elements and fill the 4th and 5th elements.
	// make sure no new backing array is created

	copy(otherPresidents[3:5], otherPresidents[:2])
	fmt.Printf("Length: %d Capacity: %d b.array %p 'otherPresidents' %q\n", len(otherPresidents), cap(otherPresidents), otherPresidents, otherPresidents)

	// copy the last 3 elements from previous slice and append 2 first elements to it.
	// make sure the backing array is still the same after append operation.

	clone := make([]string, 3, 5)
	copy(clone, otherPresidents[2:])
	fmt.Printf("Length: %d Capacity: %d b.array %p 'clone' %q\n", len(clone), cap(clone), clone, clone)

	for _, s := range otherPresidents[:2] {
		clone = append(clone, s)
	}
	fmt.Printf("Length: %d Capacity: %d b.array %p 'clone' %q\n", len(clone), cap(clone), clone, clone)

	// slice the previous slice from the 2nd to 4th element inclusively, then append an element.
	// ensure a new backing array is created
	// change the 3rd element of previous slice.
	// make sure it does not affect the content of new slice
	selected := otherPresidents[1:4:4]
	selected = append(selected, "James Madison")
	fmt.Printf("Length: %d Capacity: %d b.array %p 'selected' %q\n", len(selected), cap(selected), selected, selected)

	otherPresidents[2] = "James Monroe"
	fmt.Printf("Length: %d Capacity: %d b.array %p 'otherPresidents' %q\n", len(otherPresidents), cap(otherPresidents), otherPresidents, otherPresidents)
	fmt.Printf("Length: %d Capacity: %d b.array %p 'selected' %q\n", len(selected), cap(selected), selected, selected)

}
