package main

import "fmt"

func main() {

	//var i int // we can't redeclare var or change there types in a block
	//var i string

	a, b := 10, 20       // creates and assign values
	b, c := 200, "hello" // b gets updated and c var is created

	{
		//shadowing b from the outer scope // not recommended
		var b string = "this is an inner block"
		fmt.Println(b)
	}

	fmt.Println(a, b, c)

}
