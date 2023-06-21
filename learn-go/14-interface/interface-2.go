package main

import "fmt"

// to impl an interface we need to implement all of it methods

type walker interface {
	walk()
	hello()
}
type person struct {
	name string
}

func (p person) walk() {
	fmt.Println("i can walk", p.name)
}
func (p person) hello() {

}

// animal struct
type animal struct {
	animal string
}

func (n animal) walk() {
	fmt.Println("i can walk", n.animal)
}

func main() {

	p := person{name: "abc"}
	n := animal{animal: "lion"}
	var w walker = p
	w.walk()
	w = n
	w.walk()

}
