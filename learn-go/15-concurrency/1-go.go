package main

import (
	"fmt"
	"time"
)

// concurrency means dealing with a lot of things at once and cpu switches between the available processes
// parallelism is doing a lot of things at once

func main() {

	go hello()                  // separate line exec // new goroutine  // spinning up a goroutine
	time.Sleep(2 * time.Second) //sleeping is a non productive cpu activity so cpu will make the switch to the another go routine
	// time.sleep is not a good idea in production
}

func hello() {

	fmt.Println("hello from the hello func")
}
