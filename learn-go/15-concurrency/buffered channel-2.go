package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	c := make(chan int, 1)
	wg.Add(2)
	//go func() {
	//	defer wg.Done()
	//
	//	c <- 100
	//	c <- 200
	//	c <- 300
	//	c <- 400
	//
	//	close(c) // when channel is closed , receiver can read the remaining data but no more send are allowed after this
	//	//c <- 500 // after closing we can't send values
	//}()
	go func() {
		defer wg.Done()
		//fmt.Println(<-c)
		for v := range c {
			fmt.Println(v)
		}
	}()

	wg.Wait()

}
