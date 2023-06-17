package main

import (
	"fmt"
	"learn-go/database"
	"log"
)

func main() {

	c, ok := database.NewConfig("diwakar ", "hello ", "mysql")
	if !ok {
		log.Println("database is not up and running")
		return
	}
	fmt.Println(c)
	c.ReadData()

}
