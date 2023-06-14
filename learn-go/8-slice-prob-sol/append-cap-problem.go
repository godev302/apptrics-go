package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}
	b := x[1:4] // index:len

	b = append(b, 999)
	inspectSlice("x", x)
	inspectSlice("b", b)

}

func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println()
}
