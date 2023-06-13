package main

import (
	"fmt"
	"learn-go/sum"
)

func main() {
	//
	sum.Add()
	fmt.Println(sum.Total)
	sum.Total = 100
	fmt.Println(sum.Total)
	sum.Sub()

	//db.ReadData()
	//db.Conn = "postgres"
	//db.ReadData()
}
