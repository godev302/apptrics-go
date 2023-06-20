package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	err := openSomething("test.txt")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("no err reported")

}

var ErrFileNotFound = errors.New("some err msg")

func openSomething(fileName string) error {

	_, err := os.OpenFile(fileName, 0, 0)
	if err != nil {
		//merging two values together and trying to return err value to the caller func
		return fmt.Errorf("%v %v", err, ErrFileNotFound)
	}
	return nil
}
