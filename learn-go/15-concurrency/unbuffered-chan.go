package main

import (
	"fmt"
	"sync"
	"time"
)

// A send on an unbuffered channel can proceed if a receiver is ready.
// unbuffered chan is giving us guarantee that value would always be received by someone
// problem is it would add latency to code because if receiver take time to get ready then my sender go routine would
// block
func main() {
	wg := &sync.WaitGroup{}

	c := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second) // adding latency to the code
		fmt.Println(<-c)
	}()

	go func() {
		defer wg.Done()
		c <- 200                    // send would not proceed until there a receiver to receive the values
		fmt.Println("sending done") // it would be printed after 3 seconds
	}()

	fmt.Println("reached to end of the main")
	wg.Wait()

}
