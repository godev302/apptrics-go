package main

import (
	"log"
	"os"
)

func main() {

	// 4 -> read , 2-> write, 1 -> exec  //
	f, err := os.OpenFile("test.txt", os.O_RDONLY, 0777)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close() // it guarantees that file would be closed when the function ends

	// do work with your file
	f.Read(nil)

	// panic happened

	//file would be closed safely
}
