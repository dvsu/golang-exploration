package main

import "fmt"

func main() {

	// this section has to be fixed and simplified

	// var names [3]string = [3]string{
	// 	"Einstein" "Shepard"
	// 	"Tesla"
	// }

	// var books [5]string = [5]string{
	// 	"Kafka's Revenge",
	// 	"Stay Golden",
	// 	"",
	// 	"",
	// 	""
	// }

	// fmt.Printf("%q\n", names)
	// fmt.Printf("%q\n", books)

	names := [...]string{
		"Einstein",
		"shepard",
		"Tesla",
	}

	books := [5]string{
		"Kafka's Revenge",
		"Stay Golden",
	}

	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
}
