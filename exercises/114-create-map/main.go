package main

import "fmt"

func main() {
	// create and print maps

	// 1. lastname-phonenumber map

	phoneNumber := map[string]string{
		"Lee":   "0123456789",
		"Brown": "0135791246",
		"Adams": "0167492498",
	}

	fmt.Printf("'phoneNumber' map: %#v\n", phoneNumber)

	// 2. productID-availability map
	availability := map[int]bool{
		123798: true,
		845984: false,
		903405: true,
	}

	fmt.Printf("'availability' map: %#v\n", availability)

	// 3. multiple phoe numbers by last name
	phoneNumbersList := map[string][]string{
		"Lee":   {"0123456789", "0173485739"},
		"Brown": {"0135791246", "0123974348"},
		"Adams": {"0167492498", "0195476983"},	
	}

	fmt.Printf("'phoneNumbersList' map: %#v\n", phoneNumbersList)

	// 4. shopping basket by customer ID. Nested map contains product ID-quantity pairs
	shoppingCart := map[int]map[int]int{
		9192: {
			123798: 3,
			903405: 5,
		},
		2756: {
			123798: 1,
		},
		6789: {
			903405: 12,
		},
	}

	fmt.Printf("'shoppingCart' map: %#v\n", shoppingCart)

	// print some random key-value pairs from the maps
	q1 := "John"
	q2 := 123798
	q3 := "Brown"
	q4cid, q4pid := 2756, 123798
	
	q1p := "N/A"
	if v, e := phoneNumber[q1]; e {
		q1p = v
	}
	fmt.Printf("%s's phone number is %s\n", q1, q1p)

	a := "is"
	if !availability[q2] {
		a = "is not"
	}
	fmt.Printf("Product with ID %d %s available\n", q2, a)
	

	phone := "N/A"
	if p := phoneNumbersList[q3]; len(p) >= 2 {
		phone = p[1];
	}
	fmt.Printf("%s's second phone number is %s\n", q3, phone)

	fmt.Printf("Customer with ID %d is currently buying %d pc(s) of product with ID %d\n", q4cid, shoppingCart[q4cid][q4pid], q4pid)
}
