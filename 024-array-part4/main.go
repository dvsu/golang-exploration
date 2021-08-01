package main

import "fmt"

func main() {
	// "keyed elements" feature in Go array

	arr1 := [3]int{
		2: 10,
		0: 12,
		1: 20,
	}

	// [12, 20, 10]
	fmt.Println(arr1)

	arr2 := [5]int{
		3: 123,
	}

	// [0, 0, 0, 123, 0]
	fmt.Println(arr2)

	arr3 := [5]int{
		2: 12,
		67,
		1: 20,
	}

	// the order matters. As the unkeyed element is added right after
	// the key element, at position 2, the unkeyed element
	// is automatically assigned to position 3
	// [0, 20, 12, 67, 0]
	fmt.Println(arr3)

	arr4 := [5]int{
		1: 12,
		67,
		4: 20,
	}

	// see the difference between arr4 and arr3
	// the unkeyed element will be assigned at position (1 + 1) = 2
	// [0, 12, 67, 0, 20]
	fmt.Println(arr4)

	// using the combination ellipsis operator and keyed elements
	arr5 := [...]int{
		5: 200,
	}

	// [0, 0, 0, 0, 0, 200]
	// length = 6
	fmt.Println(arr5)

	// using the combination ellipsis operator, keyed and unkeyed elements
	// see the difference
	arr6 := [...]int{
		180,
		5: 200,
	}

	// [180, 0, 0, 0, 0, 200]
	// length = 6
	fmt.Println(arr6)

	// using the combination ellipsis operator, keyed and unkeyed elements
	// see the difference
	arr7 := [...]int{
		5: 200,
		180,
	}

	// [0, 0, 0, 0, 0, 200, 180]
	// length = 7
	fmt.Println(arr7)

	// keyed elements can be used to improve readability by replacing magic numbers
	const (
		F = iota
		NVDA
		TSLA
	)

	// it does not matter whether or not the order is shuffled, as each element
	// has been mapped to its designated index. See TSLA comes before NVDA
	// but the actual order is still maintained [F, NVDA, TSLA]
	shares := [...]float64{
		F:    13.95,
		TSLA: 687.20,
		NVDA: 194.99,
	}

	fmt.Printf("%.2f\n", shares[NVDA])
}
