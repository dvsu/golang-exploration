package main

import (
	"fmt"
	"os"
	"strconv"
)

func feetToMeter(f float64) float64 {
	return f / 3.281
}

func main() {
	// Itoa (integer-to-string) will never fail (no error handling required)
	s := strconv.Itoa(123) // convert int 123 to str "123"
	fmt.Println(s)

	// However, Atoi (string-to-integer) may fail because there is possibility that
	// the input string is empty or in alphanumeric
	// Thus the error has to be handled properly

	str := "567"

	i, err := strconv.Atoi(str)
	if err != nil {
		// when error occurs, i.e. err is not nil, do something
		fmt.Println("ERROR:", err)
		return
	}

	// execute the next lines of code if error does not occur
	fmt.Printf("Converted %q to %d successfully.\n", str, i)

	// Feet to meter conversion
	// The code snippet will cause panic if a user forgets to give an argument.
	// In the next section below, we will use simple statement to overcome this problem
	// f, err := strconv.ParseFloat(os.Args[1], 64)
	// if err != nil {
	// 	fmt.Println("ERROR:", err)
	// 	fmt.Printf("%q is not a number\n", os.Args[1])
	// 	return
	// }

	// fmt.Printf("%.3f feet = %.3f meters.\n", f, feetToMeter(f))

	// the error handling can also be simplified using simple statement
	// However, the scope of n and err is only within the curly bracket block
	if n, err := strconv.Atoi("567"); err == nil {
		fmt.Println("Everything is fine. Execute something in here")
		fmt.Printf("%[1]d Type is %[1]T\n", n)
	}

	var f float64
	// here we will fix the feet to meter conversion by handling every possible case
	// a good thing of using the simple statement is, the return keyword is not required
	// because the execution is only performed when the condition is met.
	if args := os.Args; len(args) != 2 {
		// if user gives no or more than one argument
		fmt.Println("Numeric argument is missing")
	} else if f, err := strconv.ParseFloat(args[1], 64); err != nil {
		// if user gives an argument but not a valid string
		fmt.Printf("Cannot convert %q\n", args[1])
	} else {
		//if user gives an argument and is valid
		fmt.Printf("%.3f feet = %.3f meters.\n", f, feetToMeter(f))
	}

	// print statement below will output 0 because the scope of the targeted "f" and "err"
	// are only within the block. This condition is called "shadowing"
	fmt.Println(f)

	// To prevent this from happening, both variables "f" and "err" can be declared beforehand
	// then both variables are assigned "=" with value inside the if-else statements
	// instead of redeclaration ":=". By declaring both variables outside the simple statement,
	// the variables become accessible anywhere inside the main function.
	// see the example below

	myInput := "135"

	var (
		x int
		e error
	)

	if x, e = strconv.Atoi(myInput); e != nil {
		fmt.Printf("Cannot convert %q\n", myInput)
	} else {
		x *= 3
		fmt.Printf("%q is successfully converted and tripled to %d\n", myInput, x)
	}

	// This time, the value of x is 405 instead of 0
	fmt.Println(x)

}
