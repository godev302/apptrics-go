package main

import "fmt"

func main() {

	a := []int{10, 20, 30, 40, 50}

	b := a[1:4] // index:len
	fmt.Println(b)
	b = a[2:5]
	fmt.Println(b)

	b = a[:]  // take the whole slice
	b = a[:3] // start from 0 index till the length provided
	b = a[2:] // from the 2nd index till the end

}
