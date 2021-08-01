package main

import "fmt"

func main() {
	// comparing array
	// to compare 2 arrays, both have to have the same array type
	// [2]int == [3]int // DON'T -> it will cause "invalid operation: mismatched types"
	// [3]int === [3]int64 // DON'T -> it will cause "invalid operation: mismatched types"
	// [3]int == [3]int // DO

	arr1 := [...]string{
		"a",
		"b",
		"c",
	}

	arr2 := [...]string{
		"A",
		"B",
		"C",
	}

	// print false
	fmt.Println(arr1 == arr2)

	// similarly, it is also impossible to assign array with different type
	// arr3 -> type [3]int
	// arr4 -> type [2]int
	// arr4 = arr3 (copying array of arr3 to arr4) // DON'T

	arr3 := [...]int{
		1,
		2,
	}

	arr4 := [2]int{
		10,
		20,
	}

	arr3 = arr4
	// now arr3 will be equal to [10, 20] as the new values are copied from arr4
	fmt.Println(arr3)

	// to overcome problem of copying array with different length,
	// we can use for range to loop over the source array
	arr5 := [3]string{
		"alpha",
		"bravo",
		"charlie",
	}

	var arr6 [5]string

	for i, e := range arr5 {
		arr6[i] = e
	}

	arr6[3], arr6[4] = "delta", "echo"
	// now the arr6 array contains "alpha", "bravo", "charlie", "delta", and "echo"
	fmt.Printf("%#v\n", arr6)

}
