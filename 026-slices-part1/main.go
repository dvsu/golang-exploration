package main

import (
	"fmt"
)

func main() {
	// array is a fixed collection; cannot grow or shrink
	// slice is a dynamic collection; can grow or shrink

	// slice empty value is nil, as compared with fixed array with
	// zero value at default

	// n is a slice with type []int, but the value is nil
	// as it is empty
	// Elements that are assigned to the slice have to have type of int only
	var n []int

	// length = 0
	fmt.Println(len(n))

	// comparing empty slice and array by printing both of them
	var (
		arr [5]string
		slc []string
	)
	fmt.Printf("%#v\n", arr) // [5]string{"", "", "", "", ""}
	fmt.Printf("%#v\n", slc) // []string(nil)

	// a slice can only be compared to nil value
	devices := []string{"keyboard", "mouse", "webcam"}
	gadgets := []string{"smartphone", "smartwatch"}

	// both cannot be compared, not because they have different type, but
	// slice can only be compared with nil value
	// devices == gadgets

	// a way to compare 2 slices is by using loop
	for i, d := range devices {
		if d != gadgets[i] {
			fmt.Println("Not same")
			break
		}
	}

	// this is valid
	fmt.Println(slc == nil) //  true

	// because slices do not have length, as long as the element type is the same,
	// these slices have the same type. It means value assignment is possible
	// regardless the length of the slice

	groupA := []string{"Amanda", "Barbara", "Catherine"}
	groupB := []string{"Dorris", "Elena"}

	fmt.Printf("Group B: Original %#v\n", groupB)

	// now groupB will be equal to ["Amanda", "Barbara", "Catherine"]
	groupB = groupA

	fmt.Printf("Group A: %#v\n", groupA)
	fmt.Printf("Group B: After reassigned %#v\n", groupB)

	// nil slice and empty slice are different
	// empty slice = initialized slice
	// nil slice = uninitialized slice
	// However, both length are the same

	// nil slice
	slc2 := []string{}

	fmt.Printf("slc2 is nil? %t\n", slc2 == nil) // false
	slc2 = nil

	// empty slice
	slc3 := []string{}

	// To check the content of 2 slices, it is a good idea to compare the length of both slices
	fmt.Printf("length of slc2 and slc3 is the same? %t\n", len(slc2) == len(slc3)) // true
	fmt.Printf("now slc2 is nil? %t\n", slc2 == nil)                                // false
	fmt.Printf("slc3 is nil? %t\n", slc3 == nil)                                    // false

	// remember direct slice comparison is not possible
	// fmt.Println(slc2 == slc3)
}
