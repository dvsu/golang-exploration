package main

import (
	"fmt"

	"github.com/dvsu/golang-exploration/exercises/104-limit-backing-array-sharing/api"
)

// CURRENT
//
//                           | |
//                           v v
//   api.temps     : [5 10 3 1 3 80 90]
//   main.received : [5 10 3 1 3]
//                           ^ ^ append changes the `temps`
//                               slice's backing array.
//
//
//
// EXPECTED
//
//   The corrected api package does not allow the `main()` to
//   change unreturned portion of the temps slice's backing array.
//                           |  |
//                           v  v
//   api.temps     : [5 10 3 25 45 80 90]
//   main.received : [5 10 3 1 3]
//
// ---------------------------------------------------------

func main() {

	received := api.Read(0, 3)

	received = append(received, []int{1, 3}...)

	fmt.Println("api.temps     :", api.All())
	fmt.Println("main.received :", received)
}
