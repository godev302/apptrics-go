package data

import "errors"

type User struct {
	FName string `json:"f_name"`
	LName string `json:"l_name"`
	Email string `json:"email"`
}

var users = map[uint64]User{
	123: {
		FName: "Bob",
		LName: "abc",
		Email: "bob@email.com",
	},
}

var ErrUserNotFound = errors.New("user id not found")

// FetchUser fetches the data from the map
func FetchUser(id uint64) (User, error) {
	u, ok := users[id]
	if !ok {
		return User{}, ErrUserNotFound
	}
	return u, nil
}
