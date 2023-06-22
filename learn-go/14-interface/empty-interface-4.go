package main

import "fmt"

type student struct {
	name string
}

func main() {
	//s := student{name: "raj"}
	////var a interface{} = "hello"  // empty interface
	//var a any = "hello" //any is an alias to empty interface
	//fmt.Println(a)
	//a = 10
	//fmt.Println(a)
	//a = true
	//fmt.Println(a)
	//a = s
	//fmt.Println(a)
	doSomething(10, 20)

}

func doSomething(a, b any) {
	var x, y int
	x, ok := a.(int) // type assertion, here it means we are expecting an int type from empty interface
	if !ok {
		fmt.Println("a type is not int")
		return
	}
	y, ok = b.(int)
	if !ok {
		fmt.Println("b type is not int")
		return
	}

	fmt.Println(x + y)

}
