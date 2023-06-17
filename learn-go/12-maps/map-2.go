package main

import "fmt"

type user struct {
	name string
}

func main() {
	u := make(map[int]user) //key = int , value = user
	u[101] = user{name: "john"}

	v, ok := u[101] // v = data, ok= bool

	//data is not there in the map
	if !ok {
		fmt.Println("not there")
		return
	}
	fmt.Println(v)

}
