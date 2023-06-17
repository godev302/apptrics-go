package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) updateEmail(email string) {
	u.email = email
}

type book struct {
	user           //embedding
	bookName       string
	publishingDate string
}

func (b book) show() {
	fmt.Println(b)
}

func main() {
	b := book{
		user: user{
			name:  "Jon",
			email: "jon@email.com",
		},
		bookName:       "Go",
		publishingDate: "2023",
	}

	b.updateEmail("jon@xyz.com")
	fmt.Println("directly accessing the data from the struct user using book variable", b.email)
	b.show()
}

// create a struct Movie { MovieName, DateOfRelease, Rating, User}
// Movie struct have two methods which allows you to update the movie name, print all the details
