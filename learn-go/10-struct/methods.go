package main

import "fmt"

type user struct {
	name  string
	age   int
	marks int
}

func (u user) show() { // func (receiver) methodName(args) returnTypes {
	fmt.Println("show method was called from the user struct")
	fmt.Println(u)
}

func (u *user) updateName(name string) {
	u.name = name // we have made u as a pointer receiver, which means that any change we do with it, it would
	// be reflected back
}

func main() {
	a := user{
		name:  "raj",
		age:   20,
		marks: 220,
	}

	a.show()
	a.updateName("abc")
	a.show()
}
