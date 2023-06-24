package main

import (
	"fmt"
	"sync"
	"time"
)

// cab var is a shared resource between goroutines,
// multiple go routines are trying to update the value of the cab without synchronization that is a problem
var cab = 1

func main() {
	var wg = &sync.WaitGroup{}
	var m = &sync.Mutex{}
	name := []string{"a", "b", "c", "d"}

	for _, n := range name {
		wg.Add(1)
		go bookCab(n, m, wg)
	}
	wg.Wait()
}

func bookCab(name string, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("welcome to the website", name)
	fmt.Println("some offers for you", name)

	// when a goroutine acquires a lock then another go routine can't access the critical section until the lock is not released
	//any read , write from other goroutines would not be allowed after lock is acquired
	m.Lock()
	//critical section // it is the place where you access shared resource
	if cab >= 1 {
		fmt.Println("car is available for", name)
		time.Sleep(3 * time.Second)
		fmt.Println("booking confirmed", name)
		cab--
	} else {
		fmt.Println("car is not available for", name)
	}
	m.Unlock() // releasing the lock
	fmt.Println()
}

// create a global variable to store result of the sum of two numbers and print the result
// create a func that accepts two int value and store the addition of the values in sum variable
// in main function run the sum function as a goroutine 10 times
