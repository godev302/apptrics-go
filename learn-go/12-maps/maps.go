package main

import "fmt"

func main() {
	// english to english
	// key value pair store
	//assigning the memory to the map
	namesList := make(map[int]string)

	//if the key doesn't exist than it would create the value otherwise it would update the value
	namesList[101] = "Ajay" // adding value in the map

	dictionary := map[string]string{
		"up":   "above",
		"down": "below",
	}

	dictionary["up"] = "abc" // updating a map using key

	//a := dictionary["up"]
	//delete(dictionary, "up")

	//for k, v := range dictionary {
	//	fmt.Println(k, v)
	//}
	fmt.Println(dictionary)

}
