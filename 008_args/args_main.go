package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(`=================================
OS Package & Args
=================================`)
	// var Args []string
	// go run main.go input1 input2 ...
	//   Args[0] -> path to program
	//   Args[1] -> input1
	//   Args[2] -> input2
	//   ...
	fmt.Printf("%#v\n", os.Args)
	// example output from "go run" command. Go will build the executable file at temporary path
	//                                                            VV
	// []string{"C:\\Users\\username\\AppData\\Local\\Temp\\go-build1234567890\\b001\\exe\\args_main.exe"}
	// let's try using "go build" command to build the executable at current path
	//   go build -o {file_name}
	//     -o parameter allows us to give a name to the compiled program
	//    example:
	//        go build -o args_executable.exe
	//    to run the program
	//        .\args_executable.exe

	//  alternatively on Linux or Mac
	//    example:
	//        go build -o orgs_executable
	//    to run the program
	//         ./args_executable

	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])
	fmt.Println("2nd argument:", os.Args[2])
	fmt.Println("and so on:", os.Args[3])
}
