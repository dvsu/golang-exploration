package main

import "fmt"

func main() {
	fmt.Println("naming convention. See comment section below")

	// always prevent unnecessary verbose name, i.e. name that is similar to what it is defined
	// try using these common abbreviations used in Go, widely found in Go standard library
	// var s string    		// string
	// var i int     		// index
	// var num int     		// number
	// var v string    		// value
	// var msg string 		// message
	// var fv string 		// flag value
	// var err error 		// error value
	// var args []string  	// arguments
	// var seen bool 		// has seen?
	// var parsed bool 		// parsing ok?
	// var buf []byte   	// buffer
	// var off int    		// offset
	// var op int    		// operation
	// var opRead int     	// read operation
	// var l int      		// length
	// var n int     		// number or number of
	// var m int    		// another number
	// var c int    		// capacity
	// var c int    		// character
	// var a int   			// array
	// var r rune   		// rune
	// var sep string   	// separator
	// var src int 			// source
	// var dst int 			// destination
	// var b byte			// byte
	// var b []byte			// buffer
	// var buf []byte		// buffer
	// var w io.Writer		// writer
	// var r io.Reader		// reader
	// var pos int			// position

	// for smaller scope
	// var bytesRead int // number of bytes read (not a good idea)
	// var n int // number of bytes read (better)

	// for larger scope, e.g. package scope, consider complete words
	//     package file
	//     var fileClosed bool   // declaration at package scope
	// see the name is using camel case

	// all capital letters for acronyms
	//     var localAPI string   // DO
	//     var localApi string   // DON'T

	// Do not stutter
	//     user.Name      	// DO
	//     user.UserName  	// DON'T

	// no underscores
	//     const MaxTime int     	// DO
	//     const MAX_TIME int  		// DON'T
}
