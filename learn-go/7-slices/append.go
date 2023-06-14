package main

import "fmt"

func main() {
	var a []int

	a = append(a, 10, 20, 30)
	a = append(a, 40)
	fmt.Println(a)
	fmt.Println("i=", len(a), cap(a))
}
