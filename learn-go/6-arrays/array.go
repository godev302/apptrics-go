package main

import "fmt"

// arrays size is fixed
func main() {

	var age [5]int
	age[2] = 20 // updating the value at the index
	age[3] = 23
	fmt.Println(age)

	var b = [5]int{10, 20, 30} //[10 20 30 0 0]
	//fmt.Println(b[5])

	c := [...]int{1, 2, 3, 4, 5} // use len according to the number of elems provided

	fmt.Println(b, c)
	fmt.Println(len(b))

	for i, v := range b {
		fmt.Printf("index %v ", i)
		fmt.Printf("value %v\n", v)
	}

	for _, v := range b { // _ = ignore values // here ignoring the index

		fmt.Printf("value %v\n", v)
	}

	for i := range b { // i = index
		fmt.Println(i)
	}

}
