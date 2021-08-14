package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	var (
		P  *computer
	)

	if P == nil {
		fmt.Println("Pointer value is nil")
	}

	mac := computer{brand: "apple"}
	pmac := &mac

	if *pmac == mac {
		fmt.Println("Both are equal")
		fmt.Printf("Address of pointer 'pmac' %p. Value: %p\n", &pmac, pmac)
		fmt.Printf("Address of pointer 'mac' %p. Value: %#v\n", &mac, mac)
	}

	pc := computer{brand: "microsoft"}
	ppc := &pc

	if pc != mac {
		fmt.Println("'mac' and 'pc' are not equal")
		fmt.Printf("'mac' address: %p 'pc' address: %p\n", pmac, ppc)
		fmt.Printf("'mac' value: %#v 'pc' value: %#v\n", mac, pc)
	}

	anotherMac := mac

	if anotherMac == mac {
		fmt.Println("'mac' and 'anothermac' have the same value")
	}

	fmt.Printf("'mac' address: %p 'anotherMac' address: %p\n", &mac, &anotherMac)

	fmt.Printf("Before change 'pc' address: %p 'pc' value: %#v\n", &pc, pc)

	// change microsoft to hp using function
	changeBrand(&pc, "hp")

	// confirm that it actually changes
	fmt.Printf("After change 'pc' address: %p 'pc' value: %#v\n", &pc, pc)

	// create a new computer
	fmt.Printf("Pointer of 'dell': %p\n", newComputer("dell"))
	fmt.Printf("Pointer of 'asus': %p\n", newComputer("asus"))
}

func changeBrand(com *computer, brand string) {
	fmt.Printf("Inside function. pointer address: %p, 'pc' address: %p 'pc' value: %#v\n", &com, com, com)
	com.brand = brand
}

func newComputer(brand string) *computer {
	return &computer{brand: brand}
}