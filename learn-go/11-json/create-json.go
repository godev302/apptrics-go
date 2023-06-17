package main

//https://www.devdungeon.com/content/working-files-go
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type person struct {
	FirstName string `json:"first_name"` //field tags // set the field name to the name specified for the json
	Password  string `json:"-"`          //- means ignore this field in json output
}

func main() {
	/*
		{
			"name" : "John"
		}
	*/

	p := []person{
		{
			FirstName: "Roy",
			Password:  "abc",
		},
		{
			FirstName: "John",
			Password:  "qwe",
		},
		{
			FirstName: "Bruce",
			Password:  "efg",
		},
	}

	//Marshal returns the JSON encoding of the value provided
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
		return
	}

	f, err := os.OpenFile("test.json", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	_, err = f.Write(jsonData)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(jsonData))

}

// create a slice of struct admin { firstName, lastName, password} , convert this to json , print the output on screen
