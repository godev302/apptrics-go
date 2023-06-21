package main

import "fmt"

// Polymorphism means that a piece of code changes its behavior depending on the
//concrete data it’s operating on // Tom Kurtz, Basic inventor

// interfaces are abstract type // it does not store anything of their own
// it defines how the method signature would be
type reader interface {
	read(b []byte) (int, error)
}

//when a method implemented over the struct is same as defined in the interface signature,
//interface would be automatically implemented

type file struct {
	name string
}

func (f file) read(b []byte) (int, error) {
	fmt.Println("inside file read method")
	s := "hello go devs"
	copy(b, s)
	return len(b), nil

}

type jsonObject struct {
	data string
}

func (j jsonObject) read(b []byte) (int, error) {
	s := `{name:"abc"}`
	fmt.Println("inside json read")
	copy(b, s)
	return len(s), nil
}

// fetch func accepts reader interface as an argument
// it means any type that impls this reader interface can be passed to r variable
// fetch is a polymorphic func
// fetch() will accept any type of value which implements reader interface
func fetch(r reader) {
	data := make([]byte, 50)
	r.read(data) // this read method would be called according to the concrete implementation passed to it
	fmt.Println(string(data))
	fmt.Println()
}

func main() {
	f := file{name: "abc.txt"} // concrete data
	j := jsonObject{data: "any json"}
	fetch(f)
	fetch(j)
}

// create a struct user with one method named as print()
// create a func accept(p printer) {}
// printer is an interface with one method named as print()
//call accept func from the main
