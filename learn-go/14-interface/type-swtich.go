package main

import "fmt"

type student struct {
	name string
}

func main() {
	s := student{name: "raj"}
	SomeWork(true)
	SomeWork(s)
}

func SomeWork(i any) {

	//type switch
	//it is a special switch that would match the case according to the type value passed to it
	// if value is of int type then int case would be exec and so on
	switch v := i.(type) {
	case int:
		fmt.Println("it is an int value", v)
	case string:
		fmt.Println("it is a string", v)
	case student:
		fmt.Println("it is a student type", v)
	default:
		fmt.Println("nothing matches")
	}

}
