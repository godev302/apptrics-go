package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	greet()
	fmt.Println("end of the main")
}

func greet() {

	// os.Args is a string type so no matter what we pass that would be always be of string type so conversion is imp
	data := os.Args[1:] // starts from the first index till the end

	if len(data) < 3 {
		log.Println("please provide name, age and marks")

		return // stops the exec of the current func // note :- it doesn't stop the program
		//os.Exit(1) // stop the program
	}

	//fmt.Println(data, "recv all the values")

	//var err error // default val is nil // it indicates no error

	name := data[0]
	ageString := data[1]
	marksString := data[2]

	age, err := strconv.Atoi(ageString)
	if err != nil {
		log.Println(err)
		fmt.Println(age) // age is set to its default when err occurs
		return           // stops the exec of the current func
	}

	marks, err := strconv.Atoi(marksString)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(name, age, marks)

}
