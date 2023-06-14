package main

import "fmt"

func main() {
	// slice points to an array (backing array , underlying array) in the memory
	var i []int // nil slice
	fmt.Printf("%#v\n", i)

	// this will cause panic as length is not available to store the value
	//i[2] = 100 // this is used to update values , not to add in slice

	if i == nil {
		fmt.Println("it is a nil slice")
	}

}
