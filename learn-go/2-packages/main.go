package main

import "learn-go/database"

func main() {
	//
	//sum.Add()
	//fmt.Println(sum.Total)
	//sum.Total = 100
	//fmt.Println(sum.Total)
	////sum.Sub()

	database.ReadData()
	database.Conn = "postgres"
	database.ReadData()
}
