package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 8, 9}

	// basic slice expression
	nums2 := nums1[:2]
	// full slice expression [start:end:cap]
	nums3 := nums1[:2:2]
	fmt.Printf("'nums1' ptr (backing array): %p cap: %d %d\n", nums1, cap(nums1), nums1)
	fmt.Printf("'nums2' ptr (backing array): %p cap: %d %d\n", nums2, cap(nums2), nums2)
	fmt.Printf("'nums3' ptr (backing array): %p cap: %d %d\n", nums3, cap(nums3), nums3)

	fmt.Println("With full slice expression")
	nums3 = append(nums3, 11, 13)
	fmt.Printf("Recheck 'nums1' ptr (backing array): %p %d\n", nums1, nums1)
	// how does the 'num3' backing array address change even though it still has a capacity of 4?
	//   if '11' and '13' are not appended to 'nums3', the maximum capacity is actually 2
	//   so appending 2 more elements will require more capacity. Thus, Go creates a new
	//   backing array
	fmt.Printf("Recheck 'nums3' ptr (backing array): %p %d\n", nums3, nums3)

	// however, if the same process is done without full slice expression, the '15'
	// will overwrite the previous value '9' and add '17' to the next in the new backing array
	fmt.Println("Without full slice expression")
	nums2 = append(nums2, 15, 17)
	// we still expect the 'nums1' slice retains the old element [1, 2, 8, 9]
	// but nums2 overwrites it because the capacity value is at max by default
	fmt.Println("'nums2' modified 'nums1's backing array. Now 'nums1' value is [1, 2, 15, 17]")
	fmt.Printf("Recheck 'nums1' ptr (backing array): %p %d\n", nums1, nums1)
	fmt.Printf("Recheck 'nums2' ptr (backing array): %p %d\n", nums2, nums2)
	// nums3
	fmt.Printf("Recheck 'nums3' ptr (backing array): %p %d\n", nums3, nums3)

	// therefore, full slice expression is very useful if user does not want to overwrite
	// the original backing array, although it can also be achieved sometimes
	// without using full slice expression. Compare 'nums4' (without) and 'nums5' (with)

	// here are the more concise solutions to 'nums3'. Let's name them 'nums4' and 'nums5'
	// 'nums4' does not overwrite 'nums1' because 'nums1' capacity is already full
	// So, any append will eventually create a new backing array
	nums4 := append(nums1[2:4], 39, 41)
	nums5 := append(nums1[:2:2], 23, 37)
	fmt.Printf("'nums1' ptr (backing array): %p cap: %d %d\n", nums1, cap(nums1), nums1)
	fmt.Printf("'nums4' ptr (backing array): %p cap: %d %d\n", nums4, cap(nums4), nums4)
	fmt.Printf("'nums5' ptr (backing array): %p cap: %d %d\n", nums5, cap(nums5), nums5)

}
