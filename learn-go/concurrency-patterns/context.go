package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)

	ctx := context.Background() // empty context where we will put timeouts,values in the future

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel() // clean resources taken up by context
	go func() {

		time.Sleep(4 * time.Second)
		c1 <- "one" // send

	}()
	PrintData(ctx, c1)

}

func PrintData(ctx context.Context, c chan string) {
	select {
	case <-ctx.Done(): // it will tell when a ctx is cancelled or finished // or time is up
		fmt.Println("timeout")
		fmt.Println(ctx.Err())
		return
	case x := <-c:
		fmt.Println(x)

	}
}

// from the main func send 3 names in a channel, use context to have timeout of 3 seconds
// receive the value from the channel using select
// for i:=1; i<=3;i++ {
//	select {}
//}
