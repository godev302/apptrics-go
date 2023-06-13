package main

import "fmt"

func main() {

	//for {
	//	//infinite loop
	//	fmt.Println("i will run forever")
	//}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	x := 0
	for x <= 5 {
		fmt.Println(x)
		x++
	}

}
