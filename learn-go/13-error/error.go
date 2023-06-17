package main

import (
	"errors"
	"fmt"
	"log"
)

var user = make(map[int]string)

func main() {
	data, err := FetchRecord(100)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(data)

}

// creating custom error , // prefix your err variables with Err word

var ErrRecordNotFound = errors.New("email not found")

func FetchRecord(id int) (string, error) {

	email, ok := user[id]
	if !ok {
		return "", ErrRecordNotFound
	}

	return email, nil
}

// create a func, that returns an error named as ErrFileNotFound,
