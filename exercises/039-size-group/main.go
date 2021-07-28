package main

import "fmt"

func sizeGroup(s int) string {
	if s >= 200 {
		return "Big"
	} else {
		return "Small"
	}
}

func objectCategory(sph bool) string {
	if sph {
		return "spherical"
	} else {
		return "non-spherical"
	}
}

func main() {
	isSphere, radius := true, 100

	var big bool

	// conditional statement, the naive way
	if radius >= 200 {
		big = true
	} else {
		big = false
	}

	if isSphere {
		if big {
			fmt.Println("Big spherical object")
		} else {
			fmt.Println("Small spherical object")
		}
	} else {
		if big {
			fmt.Println("Big non-spherical object")
		} else {
			fmt.Println("Small non-spherical object")
		}
	}

	// a more concised solution
	fmt.Printf("%s %s object\n", sizeGroup(radius), objectCategory(isSphere))
}
