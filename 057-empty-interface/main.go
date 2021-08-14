package main

import "fmt"

func main() {
	// we can store any type of value in an empty interface value

	// To declare an empty interface
	var (
		empty interface{}
		ifc interface{}
		bulk []interface{}
	)	
	empty = []int{11, 22, 33}
	fmt.Println(empty)

	empty = "change it to string"
	fmt.Println(empty)

	empty = 3.14
	fmt.Println(empty)

	// But we can't directly use the dynamic value
	// of an empty interface value

	// this does not work because the of mismatched types
	// empty is an empty interface type but 12 is an int
	// empty = empty * 12 

	// first, we have to extract the dynamic type and value from it.
	// Currently, it has a dynamic type of int and dynamic value of 3.14

	// Again, we can use type assertion to extract the number
	f, ok := empty.(float64)
	if ok {
		fmt.Printf("The dynamic value of the empty interface is %.2f\n", f)
	}

	// we can also perform operation in line
	ff := empty.(float64) * 1.23456
	fmt.Printf("The final value of 'ff' is %.5f\n", ff)

	// what if the empty interface has a dynamic type of slice?
	ifc = []int{10, 20, 30, 40, 50}

	// this works
	fmt.Println(ifc)

	// but this does not, because the len function should be applied to the
	// dynamic type slice, not the empty interface.
	// fmt.Println(len(ifc))

	// The workaround is to use type assertion
	l := len(ifc.([]int))
	fmt.Printf("The length of the slice is %d\n", l)

	// What about assigning value to an empty interface slice?
	// This does not work, because the variable is a slice, not a
	// single value
	// bulk = []int{10, 20}

	// We can use for-range loop to fill the empty interface slice
	source := []int{90, 80, 70}

	for _, v := range source {
		bulk = append(bulk, v)
	}

	// now bulk contains [90, 80, 70]
	fmt.Println(bulk)

	// to extract
	for _, i := range bulk {
		if d, ok := i.(int); ok {
			fmt.Printf("%d extracted from empty interface\n", d)
		}
	}
	
	// Go is designed to be a type-safe programming language.
	// Although empty interface gives type flexibility,
	// it may also introduce bugs which potentially caused
	// by type mismatch. Therefore, use empty interface
	// only if necessary.
}
