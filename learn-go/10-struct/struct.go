package main

import "fmt"

type user struct { // user is a new data type
	name  string // fields
	age   int    //struct fields would be set to there default if value is not provided
	marks int
}

func main() {

	//var u user
	//u.name = "raj"
	//u.age = 22
	//u.marks = 80
	u := user{
		name:  "raj",
		age:   22,
		marks: 220,
	}
	fmt.Println(u)
	fmt.Println(u.name)

	// create a struct for an employee, fields :- empId, empName, empAddress

}
