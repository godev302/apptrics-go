// Package sum is a library package
package sum

import "fmt"

var Total int

// making first letter uppercase of a func or a variable exports it to the other packages

func Add() {
	a, b := 10, 20
	Total = a + b
	print()
	//fmt.Fprintf()
}

func print() {
	fmt.Println(Total)
}
