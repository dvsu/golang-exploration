package main

import "fmt"

type device struct {
	maker        string
	yearReleased int
	weight       float64
	weightUnit   string
}

type smartphone struct {
	device         // this is called anonymous field
	os             string
	screenSize     float64
	screenSizeUnit string
	weight       float64
}

func main() {
	// OOP in Go using struct
	// Go does not use inheritance
	// Instead, it embeds the child/ inner struct within parent struct.
	// (composition (has-a relationship) in OOP)
	// That's why this technique is called "embedding" in Go.
	pixel4 := smartphone{
		device: device{
			maker:        "Google",
			yearReleased: 2019,
			weight:       162,
			weightUnit:   "gram",
		},
		os:             "Android",
		screenSize:     5.7,
		screenSizeUnit: "inch",
	}

	// the advantage of using anonymous field is that we can access the field
	// value of inner struct as if it's outer struct's field
	// pixel4.maker is equvalent to pixel4.device.maker

	fmt.Printf("pixel4.device.maker: %s\n", pixel4.device.maker)
	fmt.Printf("pixel4.maker: %s\n", pixel4.maker)

	// But if both inner and outer structs have the same field name, the
	// field of outer/ parent struct takes priority.
	// See what will happen if we print the value of "weight" field

	// it is zero because the pixel4.weight has higher priority than pixel4.device.weight
	fmt.Printf("pixel4.weight: %f\n", pixel4.weight)

	// To get the value of "weight" from inner struct, we have to use pixel4.device.weight
	fmt.Printf("pixel4.device.weight: %f\n", pixel4.device.weight)

}