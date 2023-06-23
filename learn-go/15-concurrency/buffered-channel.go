package main

import "fmt"

// A send on a buffered channel can proceed if there is room in the buffer
// We have to manually make sure if all the values are received
func main() {
	c := make(chan int, 2) // [100,300]
	c <- 100
	c <- 200
	fmt.Println(<-c) // receive from the channel would empty one value from the buffer
	c <- 300
	fmt.Println("all the signals are sent")
}
