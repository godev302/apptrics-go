package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type user struct {
	FirstName string `json:"first_name"`
}

func main() {
	jsonData, err := os.ReadFile("test.json")
	if err != nil {
		log.Println(err)
		return
	}
	var u []user
	err = json.Unmarshal(jsonData, &u)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(u)
	fmt.Println(string(jsonData))

}
