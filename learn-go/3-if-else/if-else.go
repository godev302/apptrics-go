package main

import "fmt"

func main() {

	a := 10
	b := 20
	c := 30

	if a > b && a > c {
		fmt.Println("a is greater")
	} else if b > a && b > c {
		fmt.Println("b is greater")
	} else {
		fmt.Println("c is greater")
	}

}
