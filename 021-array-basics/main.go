package main

import "fmt"

func main() {
	var companies [6]string

	// see the array print output by using various verbs
	fmt.Printf("Companies: %T\n", companies)  // type of array
	fmt.Println("Companies:", companies)      // array with its zero values
	fmt.Printf("Companies: %q\n", companies)  // array with its zero values quoted
	fmt.Printf("Companies: %#v\n", companies) // array with its zero values quoted + comma + array type

	name := "The Coca-Cola"

	// add array element using index expression
	companies[0] = "Amazon"
	companies[1] = "Apple"
	companies[2] = name + " Company" // this is valid
	companies[3] = "Facebook"
	companies[4] = "Google"
	companies[5] = "Microsoft"

	fmt.Printf("Companies: %#v\n", companies)

	for _, c := range companies {
		fmt.Println(c)
	}

	// another way to initialize array elements
	weather := [6]string{
		"sunny",
		"cloudy",
		"rain",
		"thunderstorms",
		"snow",
		"fog",
	}

	fmt.Printf("%q\n", weather)

	// use ellipsis operator to determine the length of array
	languages := [...]string{
		"Java",
		"Python",
		"JavaScript",
		"Go",
		"TypeScript",
		"Dart",
		"Kotlin",
		"C++",
		"C",
		"C#",
	}

	// the length of languages array is equal to 10
	fmt.Printf("%#v\n", languages)
}
