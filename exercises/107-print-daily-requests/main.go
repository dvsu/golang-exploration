package main

import "fmt"

func main() {
	// There are 3 request totals per day (8-hourly)
	const N = 3

	// DAILY REQUESTS DATA (8-HOURLY TOTALS PER DAY)
	reqs := []int{
		500, 600, 250, // 1st day: 1350 requests
		200, 400, 50, // ...
		900, 800, 600,
		750, 250, 100,
		150, 654, 235,
		320, 534, 765,
		121, 876, 285,
		543, 642,
		// the last element is missing (your code should be able to handle this)
		// that is why you shouldn't calculate the `size` below manually.
	}

	// calculate the daily and grand totals

	size := (len(reqs) / N)

	if len(reqs)%N != 0 {
		size += 1
	}

	daily := make([][]int, 0, size)

	// first approach
	for i := 0; i < cap(daily); i++ {
		if (i*N)+N > len(reqs) {
			daily = append(daily, reqs[(i*N):])
			continue
		}
		daily = append(daily, reqs[(i*N):(i*N)+N])
	}
	// first approach ends here

	/*
		// alternative
		for N <= len(reqs) {
			daily = append(daily, reqs[:N])
			reqs = reqs[N:]  // move the pointer
		}

		// handle the remaining elements, if the length of is less than 3
		if len(reqs) > 0 {
			daily = append(daily, reqs)
		}
		// alternative ends here
	*/

	// calculate daily total
	var grand int
	for i, row := range daily {
		var sum int
		for _, req := range row {
			sum += req
			grand += req
		}
		fmt.Printf("Day %d: %d requests\n", i+1, sum)
	}
	fmt.Printf("Grand total: %d requests\n", grand)
}
