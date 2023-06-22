package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	c := make(chan int) // unbuffered channel // it is a channel without any size
	wg.Add(2)
	go add(10, 20, c, wg)

	go func() {
		defer wg.Done()
		// it will block the current go routine
		// until it will not receive the values
		res := <-c // chan recv is a blocking call
		fmt.Println("result is ready")
		fmt.Println(res)

	}()
	wg.Wait()

}

func add(a, b int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := a + b
	time.Sleep(3 * time.Second)
	c <- sum // send signal
}
