package main

import "fmt"

var aaa int = 123

var aa = 42

var b bool

// creating custom data type. New data type 'customtype' will be equivalent to 'int'
type customtype int

var z customtype

func main() {

	fmt.Println(`=================================
Println, Printf & Sprintf
=================================`)

	fmt.Println("First line")
	foo()
	fmt.Println("Additional line")
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("Last one")

	fmt.Println(aaa)

	fmt.Println(aa)

	// a var variable with no value will always have its falsy value as its initial value
	fmt.Println(b)

	fmt.Println("Checking the value of variable z: ", z)
	z = 122222
	fmt.Println("Assigning new value to variable z: ", z)
	fmt.Printf("Type of z: %T\n", z)
	fmt.Println(`
However, value of new variable with type 'int' cannot be assigned as type 'customtype', 
	e.g. v := z. 
Doing so will cause an error.
Alternatively, the data type can be converted to make the assignment possible`)

	v := int(z)

	fmt.Printf("After conversion to %T, %d\n", v, v)

	x := 10
	y := false
	z := 3.14
	// %q -> value with quote included
	fmt.Printf("The text is %q\n", "Computer")
	// if a number is given for %q verb, it will print a unicode character
	fmt.Printf("Comparing various verbs and values %q %q %s\n", 77, "77", "77")
	fmt.Printf("Decimal %d\n", x)      // %d -> integer value
	fmt.Printf("Hexadecimal %#x\n", x) // %#x -> hexadecimal value with leading zero; lower case characters, including the 'x'
	fmt.Printf("Hexadecimal %#X\n", x) // %#x -> hexadecimal value with leading zero; upper case characters, including the 'X'
	fmt.Printf("Hexadecimal %x\n", x)  // %x -> hexadecimal value without leading zero; lower case characters
	fmt.Printf("Hexadecimal %X\n", x)  // %X -> hexadecimal value without leading zero; upper case characters
	fmt.Printf("Binary %#b\n", x)      // %d -> binary value with leading zero '0b'
	fmt.Printf("Binary %b\n", x)       // %d -> binary value without leading zero
	fmt.Printf("Boolean %t\n", y)      // %t -> boolean value
	fmt.Printf("Float %f\n", z)        // %f -> float value
	fmt.Printf("Float %#x\n", x)
	// Sprintf allows formatted string to be reassigned as a new string variable
	text := fmt.Sprintf("new text is here %#b", x)
	fmt.Println(text)

	fmt.Printf("variable 'text' type: %T\n", text) // %T -> variable type

	// %v is a general verb which can be used to show any value
	capital := "Rome"
	country := "Italy"
	population := 4.257 // millions
	year := 2019

	fmt.Printf("The capital city of %v is %v. The population was %v millions in %v\n", country, capital, population, year)

	// In general, the number of verbs must match to the number of variables after comma.
	// However, using an argument index allows extra flexibility as shown below
	firstName := "Adam"
	lastName := "Smith"
	fmt.Printf("Smith: Hi, I'm %v %v.\nSouthgate: Hi Mr %[2]v! Welcome to our annual event!\n", firstName, lastName)

	// verb can also be used as type check. It means, it will never allow variables that do not match the type of verb
	// fmt.Printf("Float %f\n", x) // It will print "Float %!f(int=10)" because x is an int, while the verb is float.
	// in fact, the %v perform the verb replacement to the right verb type under the hood

	// However, using the same type gives some advantages. In the case of float, it gives higher precision
	// see the difference
	fmt.Printf("Comparing precision %[1]f %[1]v\n", population)
	// The precision of float can also be adjusted. Rounding will be automatically performed to match the precision
	fmt.Printf("Adjusting precision %.0[1]f %.1[1]f %.2[1]f %.7[1]f\n", population)
}

func foo() {
	fmt.Println("Inside foo")
}
