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
	f, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("ERROR:", err)
		fmt.Printf("%q is not a number\n", os.Args[1])
		return
	}

	fmt.Printf("%.3f feet = %.3f meters.\n", f, feetToMeter(f))
}
