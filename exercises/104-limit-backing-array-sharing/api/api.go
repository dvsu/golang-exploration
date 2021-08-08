package api

var temps = []int{5, 10, 3, 25, 45, 80, 90}

// Read returns a slice of elements from the temps slice.
func Read(start, stop int) []int {

	// portion := temps[start:stop] // this line does not work
	// because the default capacity is 7

	// Solution and explanation:
	// Without using full slice expression, it will still
	// assume the original capacity. Thus, appending an element
	// or two may not create a new backing array, as the
	// capacity is still enough. At the same time, it will
	// also rewrite the old values, which we want to avoid.
	// The simple workaround is to use full slice expression
	// and limit the capacity = length.
	// As the capacity is at max, appending a new element
	// will create a new backing array, but does not rewrite
	// the original backing array.
	portion := temps[start : stop : stop-start]

	return portion
}

// All returns the temps slice
func All() []int {
	return temps
}
