package db

import "fmt"

var Conn = "mysql"

func ReadData() {
	fmt.Println("reading data from", Conn)
}
