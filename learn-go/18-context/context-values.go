package main

import (
	"context"
	"fmt"
)

// Key is a new datatype which is based on string type
type Key string

func main() {
	ctx := context.Background()
	const k Key = "anyKey"
	ctx = context.WithValue(ctx, k, "abc") // putting value in the context associated with a key
	FetchValue(ctx, k)
}

func FetchValue(ctx context.Context, k Key) {
	v := ctx.Value(k)
	s, ok := v.(string)
	if !ok {
		fmt.Println("value is not there or of a different type")
		return
	}
	fmt.Println(s)

}

// create a new key with value stringKey
// put 100 (int) in the context value and fetch it in function
