package main

import (
	"fmt"
	"unsafe"
)

type StringHeader struct {
	pointer uintptr
	length int
}

func main() {
	// Similar to slice header, string header also contains a few properties
	// such as pointer and length, except capacity because
	// string value is read-only byte slice and it cannot grow in size.

	// String is a read-only byte slice because any identical string value
	// and its slices share the same backing array. If modification
	// were possible, it would have break the whole variables that
	// share the same backing array

	// However, byte slice can be modified because it creates a new backing array.
	// It means modification is possible without affecting the original backing array.
	// The capacity also can be larger than its length.

	ocean := "Atlantic"
	oceanBytes := []byte(ocean)

	oceanBytes[0] = '@'
	oceanBytes[1] = '-'
	fmt.Printf("%s\n",oceanBytes)

	// we will prove that identical string values will share and point to the same backing address
	ptr1 := *(*StringHeader)(unsafe.Pointer(&ocean))
	fmt.Printf("'ocean' %+v\n", ptr1)

	anotherOcean := "Atlantic"
	ptr2 := *(*StringHeader)(unsafe.Pointer(&anotherOcean))
	fmt.Printf("'anotherOcean' %+v\n", ptr2)

	// its slice also shares the same backing array
	oceanSliced := ocean[:3]
	ptr3 := *(*StringHeader)(unsafe.Pointer(&oceanSliced))
	fmt.Printf("'oceanSliced' %+v\n", ptr3)

	// but a character difference will create a new backing array
	yetAnotherOcean := "Atlantico"
	ptr4 := *(*StringHeader)(unsafe.Pointer(&yetAnotherOcean))
	fmt.Printf("'yetAnotherOcean' %+v\n", ptr4)	

	// 		  backing array #1        backing array #2             backing array #3
	//			   V                        V                            V
	//          string ------------------> byte ---------------------> string
	//			            allocate                   allocate
	//		           new backing array          new backing array

	// backing array #1 != backing array #3
	// therefore, only perform conversion when necessary
}