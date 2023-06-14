package main

import (
	"fmt"
	"os"
)

func main() {
	//it would print the args passed from the terminal,
	//0 index is always location to exec
	fmt.Println(os.Args)

	list := os.Args[1:]
	fmt.Println(list)

}
