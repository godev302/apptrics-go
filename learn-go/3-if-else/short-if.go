package main

import "fmt"

func main() {

	if a, b := 10, 20; a == b {
		fmt.Println("a and b are equal")
	} else if a > b {
		fmt.Println("a is greater than b ")
	} else {
		fmt.Println("not sure")
	} // life of var declared inside the if block ends with if

	//fmt.Println(a) // would not work

}
