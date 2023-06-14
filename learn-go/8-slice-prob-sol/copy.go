package main

import "fmt"

func main() {
	a := []int{900, 670, 610, 950, 870, 200}

	// allocating a new backing array for slice b
	b := make([]int, len(a)) // len is important when we want to copy elems

	copy(b, a)
	b[0] = 9999
	fmt.Println(a)
	fmt.Println(b)

}
