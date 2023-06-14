package main

import (
	"fmt"
	"log"
)

func main() {

	msg, ok := hello("ajay", 20, 100)
	if !ok { // ok == false
		log.Println("process failed", msg)
		return
	}
	fmt.Println(msg)

}

func hello(name string, age int, marks int) (string, bool) {
	if name == "" {
		return "please provide a name", false // it will stop the current func and return values
	}
	if age == 0 {
		return "please provide your age", false
	}

	if marks == 0 {
		return "please provide your marks", false
	}

	fmt.Println(name, age, marks)
	return "success", true

}

/*

create a func that accepts two number, and return sum of the numbers


*/
