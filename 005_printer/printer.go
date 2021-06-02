package main

import "fmt"


var aaa int = 123

var aa = 42

var b bool

// creating custom data type. New data type 'customtype' will be equivalent to 'int'
type customtype int

var z customtype

func main() {

	fmt.Println("First line")
	foo()
	fmt.Println("Additional line")
	for i := 0; i < 30; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("Last one")

	fmt.Println(aaa)

	fmt.Println(aa)

	// a var variable with no value will always have its falsy value as its initial value
	fmt.Println(b)

	fmt.Println("Checking the value and type of variable z")
	fmt.Println(z)
	fmt.Println("Assigning new value to variable z")
	z = 122222
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(`
	However, value of new variable with type 'int' cannot be assigned as type 'customtype', 
		e.g. v := z. 
	Doing so will cause an error.
	Alternatively, the data type can be converted to make the assignment possible`)

	v := int(z)

	fmt.Printf("After conversion to %T, %d\n", v, v)

	x := 10
	fmt.Printf("%d\n", x)
	fmt.Printf("%#x\n", x)
	fmt.Printf("%X\n", x)
	fmt.Printf("%#b\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%b\n", x)

	// Sprintf allows formatted string to be reassigned as a new string variable
	text := fmt.Sprintf("new text is here %#b", x)

	fmt.Println(text)
	fmt.Printf("%T", text)
}

func foo() {
	fmt.Println("Inside foo")
}
