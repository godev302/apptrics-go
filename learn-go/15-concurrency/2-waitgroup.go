package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	//wg.Add(10) // counter to add number of goroutines that we are starting or spinning up
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go work(i, wg) //each func call would create a goroutine // 10 goroutines
	}

	wg.Wait() // it will wait until counter resets to zero
	fmt.Println("end of main")
}

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrements the counter by one // defer make sure counter is always decremented even in case of panic
	fmt.Println("i am doing some work", id)

}

func sum(a, b int) {
	fmt.Println(a + b)
}
