package main

import "fmt"

func main() {
	y := 10

	var p *int // this var can store address of var of int type // default is nil

	// any changes made by p would be reflected back to the var y because p stores the memory address of y var and it is directly manipulating it
	p = &y // storing the address of y in p pointer

	y = 20
	fmt.Println(y)
	*p = 40 // * is a dereference operator to access the memory stored in the pointer
	fmt.Println(y)
	fmt.Println(*p)

	fmt.Println(&y)
	fmt.Println(p)
}
