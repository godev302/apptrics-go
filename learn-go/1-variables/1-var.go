package main

import (
	"fmt"
)

func main() {

	var s string // default value is empty string of string datatype ("")
	var i int    // 0 is default value

	fmt.Println(s, i, "hello")

	fmt.Printf("%q\n", s) // \n for new line

	//var name = "Bob"
	////name = 1 // go is a statically compiled language which means it need to know the type at the compile time
	//fmt.Println(name)
	//
	//age, marks, check := 33, 99.7, true //shorthand
	//fmt.Println(age, marks, check)
	//

	var (
		name  string = "Raj"
		age          = 33
		marks int
	)
	fmt.Println(name, age, marks)

	fmt.Println("end of the main")

	//http.StatusOK
	//time.Second
	//os.O_APPEND

}
