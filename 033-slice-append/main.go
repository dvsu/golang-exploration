package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 7}

	fmt.Printf("Step 0: ptr (header): %p ptr (backing array): %p Length %d Capacity %d Values: %d\n", &nums, nums, len(nums), cap(nums), nums)

	// How 'append' works
	// Behind the scene, Go will create a new backing array when we try to append to full-capacity array
	// The backing array comes with extra capacity to anticipate future append.
	// Then, the element values from old backing array will be copied to the new backing array.
	// The extra capacity will be filled with zero value.

	// Notes:
	// 1. Creating new backing array will also change the existing backing array address to new one.
	// 2. The extra capacity is created because 'copying' is an expensive operation.
	//    By having extra capacity, Go doesn't need to create a new backing array every time we append an element
	// 3. The new backing array is usually created when length < capacity / 2

	// Sequence:
	// - Declare empty slice. length: 0 capacity: 0
	// - append an element, length 0 -> 1
	// - Create new backing array. length: 1 capacity: 1
	// - append an element, length 1 -> 2
	// - Create new backing array. length: 2 capacity: 2
	// - append an element, length 2 -> 3
	// - Create new backing array. length: 3 capacity: 4
	// - append an element, length 3 -> 4
	// - append an element, length 4 -> 5
	// - Create new backing array. length: 5 capacity: 8
	// - append an element, length 5 -> 6
	// - append an element, length 6 -> 7
	// - append an element, length 7 -> 8
	// - append an element, length 8 -> 9
	// - Create new backing array. length: 9 capacity: 16
	// - append an element, length 9 -> 10
	// - append an element, length 10 -> 11
	// - append an element, length 11 -> 12
	// - append an element, length 12 -> 13
	// - append an element, length 13 -> 14
	// - append an element, length 14 -> 15
	// ...

	// when the number of element is small enough, the growth ratio of the capacity is equal to 2.
	// As the capacity hits 1024 elements, the ratio drops to around 1.25 - 1.3x

	// append 3 elements in the middle of slice [1, 3, x, x, x, 5, 7]
	// 1. duplicate the tail of slice, as many as the number of extra elements
	// nums = [1, 3, 5, 7, 3, 5, 7]
	nums = append(nums, nums[1:4]...)
	fmt.Printf("Step 1: ptr (header): %p ptr (backing array): %p Length %d Capacity %d Values: %d\n", &nums, nums, len(nums), cap(nums), nums)

	// 2. insert the extra element by appending the head of slice and the elements
	// it will automatically overwrite the previous element value
	// nums = [1, 3, 13, 15, 17, 5, 7]
	nums = append(nums[:2], 13, 15, 17)
	fmt.Printf("Step 2: ptr (header): %p ptr (backing array): %p Length %d Capacity %d Values: %d\n", &nums, nums, len(nums), cap(nums), nums)

	// 3. reveal the tail [5, 7] by reassign the whole slice
	// from [1, 3, 13, 15, 17]
	// to [1, 3, 13, 15, 17, 5, 7]
	nums = nums[:7]

	fmt.Printf("Step 3: ptr (header): %p ptr (backing array): %p Length %d Capacity %d Values: %d\n", &nums, nums, len(nums), cap(nums), nums)

}
