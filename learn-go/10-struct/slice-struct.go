package main

import "fmt"

type student struct {
	name  string
	email string
}

func (s student) print() {
	fmt.Println(s.name, s.email)
}

func main() {
	//i := []int{10, 20, 30, 40}
	s := []student{
		{ // 1 student data  // 0 index of the slice
			name:  "s1",
			email: "s1@email.com",
		},
		{ // 2nd student data  // 1st index of the slice
			name:  "s2",
			email: "s2@email.com",
		},
	}

	fmt.Println(s[0])
	fmt.Println(s[1].name, "accessing the data from the 1st index")

	for _, sData := range s {
		fmt.Println(sData.name)
		sData.print()
	}

}

// create a struct for list of movies (slice of movies struct), create one method to update movieName
