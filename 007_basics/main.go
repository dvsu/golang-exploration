package main

import (
	"fmt"
	"path"
)

// declaration at package scope
var area float32

func main() {

	fmt.Println(`=================================
Variables
=================================`)
	var dir, file string
	dir, file = path.Split("path/to/your/file.txt")
	fmt.Println(dir, file)

	// discard unnecessary values
	// this example will discard directory value
	var file2 string
	_, file2 = path.Split("journey/to/the/west.pdf")
	fmt.Println(file2)

	// combining short declaration assignment
	dir3, file3 := path.Split("/home/ubuntu/Downloads/software_v1.2.deb")
	fmt.Println(dir3)
	fmt.Println(file3)

	fmt.Println(`=================================
Declaration vs. Short Declaration
=================================`)

	// recommended to use declaration if the initial value is unknown
	var distance float32
	// as the value is unknown, the print statement below will output 0
	fmt.Println(distance)
	// Similarly, a variable declaration with type string will output
	// an empty string "" if it is printed

	// it is even better if the declaration is done at package scope. See above main function

	// group declaration for greater readability, or if variables may relate to each other
	var (
		// not really related
		pcOwner string

		// closely related, grouped together for greater readability
		cpuMaker string
		cpuPrice float32
	)
	fmt.Println(pcOwner, cpuMaker, cpuPrice) // expected output: """"0 (empty string, empty string, zero)

	// use short declaration if
	// 1. values are known
	// 2. keep code concise
	// 3. redeclaration

	// var height, weight = 178, 68 // DON'T
	height, weight := 178, 68 // DO!
	fmt.Println(height, weight)

	// DON'T
	// height = 182  // assign new value
	// gender := "male" // declare new variable

	// DO
	height, gender := 182, "male"
	fmt.Println(height, gender)

	// noted that redeclaration may cause 'shadowing'

	fmt.Println(`=================================
Type Conversion
=================================`)
	mass_density, mass, volume := 1, 1000, 0.879

	// mass_density = mass / volume // it doesn't work because types are different
	// conversion is required. In this case, we will convert to int

	// mass_density = mass / int(volume) // this also doesn't work because int of 0.879 is 0 -> destructive operation
	// it will cayse division by zero error
	// to prove it
	fmt.Println(volume, int(volume))

	// here is the correct solution
	// first, convert to float64 then the final product is converted to int to match mass_density data type
	mass_density = int(float64(mass) / volume)
	fmt.Println(mass_density, mass, float64(mass))

}
