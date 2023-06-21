package main

import "fmt"

type Runner interface {
	Run()
}
type Walker interface {
	Walk()
}

type WalkRunner interface {
	Runner // interface embedding
	Walker
	// walk()
	//run()
}

// Human struct impl all the three interfaces
type Human struct{ name string }

func (h Human) Walk() {
}
func (h Human) Run() {
}

type Robo struct{ name string }

func (r Robo) Run() {}
func (r Robo) Code() {
	fmt.Println("robo is coding")
}

func main() {

	h := Human{name: "John"}
	robot := Robo{name: "r1"}
	_, _ = h, robot
	var r Runner = robot
	r = h

	//type assertion
	//it checks if the type is present in the interface or not
	//type assertion take the concrete value out of the interface
	//v, ok := r.(Robo) // returns struct, bool
	//if !ok {
	//	fmt.Println("you are not robo")
	//	return
	//}
	//v.Code()

}
