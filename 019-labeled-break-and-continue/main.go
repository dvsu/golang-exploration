package main

import (
	"fmt"
	"strings"
)

func main() {
	t := `At one time they got the idea that 
what the public wanted was not anything instructive but 
something light and amusing`

	l := strings.Fields(t)

	// queried words
	q := []string{"the", "and"}

	// simple query using nested loop
	// it will also return duplicated words
	fmt.Println("Regular nested for loop")
	for _, qr := range l {
		for _, word := range q {
			if qr == word {
				fmt.Printf("%q\n", qr)
			}
		}
	}

	// However, there will also be a case when
	// we want to prevent result duplication.
	// Labeled break can be useful in this case
	// because it can 'break' to predifed location
	fmt.Println("Labeled break")
queries1:
	for _, qr := range q {
		for _, word := range l {
			if qr == word {
				fmt.Printf("%q\n", qr)
				break queries1
			}
		}
	}

	// Similarly, labeled continue can also be
	// used to perform similar operation
	// the behavior below is very similar to nested for loop.
	fmt.Println("Labeled continue")
queries2:
	for _, qr := range q {
		for _, word := range l {
			if qr == word {
				fmt.Printf("%q\n", qr)
				continue queries2
			}
		}
	}

	// labeled break can also be used inside switch statement.
	// below is an example of using unlabeled break to perform
	// word filter. The result is not desirable and can be fixed
	// by using labeled break on the next block
	fmt.Println("Unlabeled break inside switch statement")
queries3:
	for _, qr := range q {
		for _, word := range l {
			switch qr {
			case "and", "or", "the":
				break
			}
			// the break, however, only breaks outside the switch statement
			// Therefore, the code below will be executed as well
			if qr == word {
				fmt.Printf("%q\n", qr)
				continue queries3
			}
		}
	}

	// below is the fix using labeled break
	fmt.Println("Labeled break inside switch statement")
queries4:
	for _, qr := range q {
	search:
		for _, word := range l {
			switch qr {
			case "and", "or", "the":
				break search
			}
			// because the queried words contain filtered words, "the" and "and"
			// the output will be empty, i.e. zero words match
			if qr == word {
				fmt.Printf("%q\n", qr)
				continue queries4
			}
		}
	}
}
