package main

import "fmt"

func main() {
	// variadic param can accept any number of values
	//fmt.Println(10, 20, 30, 40, 50, 60)
	show("10", 20, 30, 40, 50, 60)
}

func show(s string, i ...int) { // variadic param should be the last in the func signature
	fmt.Printf("%T\n", i)
	fmt.Println(i)
}
