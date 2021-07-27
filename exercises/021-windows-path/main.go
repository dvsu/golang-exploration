package main

import "fmt"

func main() {
	path := "c:\\program files\\duper super\\fun.txt\n" +
		"c:\\program files\\really\\funny.png"

	fmt.Println(path)
	// equivalent path using string literal
	pathEqv := `c:\program files\duper super\fun.txt
c:\program files\really\funny.png`

	fmt.Println(pathEqv)
}
