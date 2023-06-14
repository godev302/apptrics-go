package main

import "fmt"

func main() {
	var i []int // nil

	//creating a backing array for i slice
	i = make([]int, 0, 100) // len = 0 , cap = 100
	i = append(i, 10, 20, 30)
	inspectSlice("i", i)
}

func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println()
}
