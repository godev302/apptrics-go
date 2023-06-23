package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	//wgWorker keep track of if the go routine work is finished or not and we wil close the channel when work is done
	var wgWorker = &sync.WaitGroup{}
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	c3 := make(chan string, 1)
	done := make(chan bool, 1)

	wgWorker.Add(3)
	go func() {
		defer wgWorker.Done()

		time.Sleep(time.Second)
		c1 <- "one" // send
	}()

	go func() {
		defer wgWorker.Done()

		time.Sleep(3 * time.Second)
		c2 <- "two"

	}()

	go func() {
		defer wgWorker.Done()
		c3 <- "three"
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//waiting for go routines to finish
		wgWorker.Wait()
		fmt.Println("closing the done channel")
		//closing the channel
		close(done) // after closing, we can read from the channel but can't send new values
	}()

	//x := <-c1 // it would block until it would not receive values from the c1 channel
	//fmt.Println(x)
	//y := <-c2 // this line would block for 3 sec
	//fmt.Println(y)
	//z := <-c3 // no block // no sleep specified
	//fmt.Println(z)
	//
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			// whichever case is not blocking exec that first
			//whichever case is ready first exec that.
			case x := <-c1: // recv over the channel
				fmt.Println("", x)
			case y := <-c2:
				fmt.Println(y)
			case z := <-c3:
				fmt.Println(z)
			case <-done: // this case will exec when channel is closed
				fmt.Println("it is closed")
				return

			}
		}
	}()
	wg.Wait()
}
