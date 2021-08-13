package main

import (
	"fmt"
	"strings"
)

func main() {
	// slices that are created from a slice usually point to the same backing array.
	// Each slice header contains the address of backing address it points to.
	// It means changing slice element values inside a function will also modify the original backing array
	// as they all point to the same address.
	// So, it is very uncommon to use pointer with a slice.

	numsSlice := []int{12, 13, 14, 15, 16}
	numsSliceFixed := make([]int, 0, 8)
	numsSlice2 := []int{20, 30, 40, 50, 60}
	numsArray := [5]int{21, 31, 41, 51, 61}

	fmt.Printf("%-13s%-18s%-22s%-23s%-40s%-20s\n", "Slice/Array", "Slice/Array Name", "Slice Header Address", "Backing Array Address", "Description", "Elements")
	fmt.Printf("%s\n", strings.Repeat("=", 140))

	// However, there is also an exception when the address of backing array changes.
	// When a backing array is out of capacity, Go will create a new backing array to accommodate and anticipate additional elements.
	// When a number is appended inside addANumber() function, Go creates a new backing array.
	// So, the original backing array cannot "see" the change.
	fmt.Printf("%-13s%-18s%-22p%-23p%-40s%[4]d\n", "Slice", "numsSlice", &numsSlice, numsSlice, "Inside main(), original")
	addANumber(numsSlice)
	fmt.Printf("%-13s%-18s%-22p%-23p%-40s%[4]d\n", "Slice", "numsSlice", &numsSlice, numsSlice, "Inside main(), after func call")
	fmt.Printf("%s\n", strings.Repeat("-", 140))

	// The workaround to the previous slice problem is to create a slice with enough capacity, so
	// all slice headers will still point to the same backing array
	numsSliceFixed = append(numsSliceFixed, []int{12, 13, 14, 15, 16}...)
	fmt.Printf("%-13s%-18s%-22p%-23p%-40s%[4]d\n", "Slice", "numsSliceFixed", &numsSliceFixed, numsSliceFixed, "Inside main(), original")
	addANumberFixed(numsSliceFixed)

	// the newly added number actually exists, but hidden. To display it, the length has to be expanded up to the 6th element.
	fmt.Printf("%-13s%-18s%-22p%-23p%-40s%[4]d\n", "Slice", "numsSliceFixed", &numsSliceFixed, numsSliceFixed[:6], "Inside main(), after func call")
	fmt.Printf("%s\n", strings.Repeat("-", 140))

	// Changing element values, however, will maintain the existing backing array.
	// So, the final check after function call proves that different slice headers
	// still point to the same backing address.
	fmt.Printf("%-13s%-18s%-22p%-23p%-40s%[4]d\n", "Slice", "numsSlice2", &numsSlice2, numsSlice2, "Inside main(), original")
	doubledSlice(numsSlice2)
	fmt.Printf("%-13s%-18s%-22p%-23p%-40s%[4]d\n", "Slice", "numsSlice2", &numsSlice2, numsSlice2, "Inside main(), after func call")
	fmt.Printf("%s\n", strings.Repeat("-", 140))
	
	// On the other hand, array is inflexible. Instead of creating a new slice header
	// for every function call, the original array gets copied and the array
	// inside a function is totally different array with the same elements.
	// So, the original array will never "see" any changes happening inside a function.
	fmt.Printf("%-13s%-18s%-22s%-23p%-40s%d\n", "Array", "numsArray", "-", &numsArray, "Inside main(), original", numsArray)
	doubledArray(numsArray)
	fmt.Printf("%-13s%-18s%-22s%-23p%-40s%d\n", "Array", "numsArray", "-", &numsArray, "Inside main(), after func call", numsArray)
	
}

func addANumber(n []int) {
	fmt.Printf("%-13s%-18s%-22p%-23p%-40s%[4]d\n", "Slice", "numsSlice", &n, n, "Inside addANumber(), before append")
	n = append(n, 5)
	fmt.Printf("%-13s%-18s%-22p%-23p%-40s%[4]d\n", "Slice", "numsSlice", &n, n, "Inside addANumber(), after append")
}

func addANumberFixed(n []int) {
	fmt.Printf("%-13s%-18s%-22p%-23p%-40s%[4]d\n", "Slice", "numsSliceFixed", &n, n, "Inside addANumberFixed(), before append")
	n = append(n, 5)
	fmt.Printf("%-13s%-18s%-22p%-23p%-40s%[4]d\n", "Slice", "numsSliceFixed", &n, n, "Inside addANumberFixed(), after append")
}

func doubledSlice(n []int) {
	fmt.Printf("%-13s%-18s%-22p%-23p%-40s%[4]d\n", "Slice", "numsSlice2", &n, n, "Inside doubledSlice(), before doubled")
	for i := range n {
		n[i] *= 2
	}
	fmt.Printf("%-13s%-18s%-22p%-23p%-40s%[4]d\n", "Slice", "numsSlice2", &n, n, "Inside doubledSlice(), after doubled")
}

func doubledArray(n [5]int) {
	fmt.Printf("%-13s%-18s%-22s%-23p%-40s%d\n", "Array", "numsArray", "-", &n, "Inside doubledArray(), before doubled", n)
	for i := range n {
		n[i] *= 2
	}
	fmt.Printf("%-13s%-18s%-22s%-23p%-40s%d\n", "Array", "numsArray", "-", &n, "Inside doubledArray(), after doubled", n)
}