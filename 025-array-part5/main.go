package main

import "fmt"

func main() {
	// "named and unnamed types" in Go array

	// this is an unnamed type
	// the type is [5]int; the underlying type is also [5]int
	arr1 := [5]int{1, 6, 7, 5, 2}

	fmt.Println(arr1)

	// example of named type
	// the type is 'name'; the underlying type is [3]string
	type person [3]string
	arr2 := person{"Adam", "Bryan", "Charlie"}

	fmt.Printf("%#v\n", arr2)

	// see the difference if comparison is made between named and unnamed types
	// even with the same elements
	arr3 := [3]string{"Adam", "Bryan", "Charlie"}

	// result is 'true' because both still have the same underlying type
	fmt.Println(arr2 == arr3)

	// what if two named types are compared, even both have the same elements?
	type student [3]string
	arr4 := student{"Adam", "Bryan", "Charlie"}

	// result is 'true' because both still have the same underlying type
	// the comparison is not possible as both have different type
	// Go will not allow this, so the code below is commented
	// fmt.Println(arr2 == arr4)

	// however, there is a way to compare both array by converting either array
	// to match the other array type.
	// let say converting arr4 from 'student' type to 'person' type
	fmt.Println(person(arr4) == arr2)

	type (
		integer   int
		something [3]integer
	)

	// the code below cannot be compared because both have different underlying types
	// _ = [3]integer{} == [3]int{}

	// however, this one is valid
	_ = [3]integer{} == something{}
}
