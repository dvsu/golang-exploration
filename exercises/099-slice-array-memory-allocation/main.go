package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"
)

const size = 1e7

func main() {
	// disable garbage collector
	debug.SetGCPercent(-1)

	// run the program to see the initial memory usage.
	report("initial memory usage")
	var arr [size]int
	report("After declaring an array")
	arr2 := arr
	report("After copying an array")
	passArray(arr2)
	// convert array to slice
	slc1 := arr[:]
	slc2 := slc1[:1000]
	slc3 := slc1[1000:10000]
	report("After slicings")
	passSlice(slc3)

	fmt.Printf("arr's size: %d bytes\n", unsafe.Sizeof(arr))
	fmt.Printf("arr2's size: %d bytes\n", unsafe.Sizeof(arr2))
	fmt.Printf("slc1's size: %d bytes\n", unsafe.Sizeof(slc1))
	fmt.Printf("slc2's size: %d bytes\n", unsafe.Sizeof(slc2))
	fmt.Printf("scl3's size: %d bytes\n", unsafe.Sizeof(slc3))
}

// passes [size]int array — about 80MB!
//
// observe that passing an array to a function (or assigning it to a variable)
// affects the memory usage dramatically
func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// only passes 24-bytes of slice header
//
// observe that passing a slice doesn't affect the memory usage
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

// reports the current memory usage
// don't worry about this code
func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
