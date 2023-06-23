package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {

	// we should close the channel when all the goroutines are finished sending
	wg := &sync.WaitGroup{}
	const job = 3
	ch := make(chan string, job) // buffered channel // size = 3

	for work := 1; work <= job; work++ {
		wg.Add(1)
		// fanout pattern // spin one go routine for one job
		go func(w int) {
			defer wg.Done()

			ch <- "work " + strconv.Itoa(w)
			time.Sleep(time.Second)
		}(work)

	}
	go func() {
		wg.Wait()
		//channel would not be closed until waitgroup counter is not back to zero as it would block the func till then
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
	//for i := 1; i <= job; i++ {
	//	fmt.Println(<-ch)
	//}

}
