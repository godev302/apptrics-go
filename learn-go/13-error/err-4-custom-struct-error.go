package main

import (
	"errors"
	"fmt"
)

// QueryError is dedicated to work with error handling, it should not be used to work with normal data
type QueryError struct {
	Func  string
	Input string
	Err   error
}

// this method is compulsory to format your err msg //
func (q *QueryError) Error() string {
	//q.Err.Error() // convert the err msg to the string
	return "main." + q.Func + ": " + "input " + q.Input + ": " + q.Err.Error()
}

var ErrNotFound = errors.New("not found")
var ErrMismatch = errors.New("mismatch")

func SearchSomething(s string) error {
	//do searching and if that is not found then return the error below
	return &QueryError{
		Func:  "SearchSomething",
		Input: s,
		Err:   ErrNotFound,
	}

}

func SearchName(name string) error {
	//do searching and if that is not found then return the error below
	return &QueryError{
		Func:  "SearchName",
		Input: name,
		Err:   ErrMismatch,
	}
}

func main() {
	var qe *QueryError
	err := SearchSomething("some data")
	if err != nil {
		//it checks the error chain to find if the custom error struct was added in the chain or not
		if errors.As(err, &qe) {
			fmt.Println("custom struct found in the chain")
			fmt.Println(qe.Func)
			return
		}
		fmt.Println("custom struct not found in the chain")
		return
	}
	fmt.Println(err)
	err = SearchSomething("random data")
	fmt.Println(err)
	err = SearchName("sahil")
	fmt.Println(err)

}

// create an error struct with two field named as {Path string , Err error}
// create a func that accepts the file path, assume that file is not found, return the error from the func as struct

/*
// run this to see dynamic err msg
_, err := strconv.Atoi("abc")
	fmt.Println(err)
	_, err = strconv.Atoi("xyz")
	fmt.Println(err)
	_, err = strconv.ParseInt("hello", 10, 64)
	fmt.Println(err)
	_, err = strconv.ParseUint("bye", 10, 64)
	fmt.Println(err)
*/
