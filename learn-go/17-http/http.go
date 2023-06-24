package main

import (
	"fmt"
	"log"
	"net/http"
)

// curl localhost:8080/home
func main() {
	http.HandleFunc("/home", HomePage)

	// this will start the server //
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

// HomePage
// http.ResponseWriter is used to write response back to the client ,
// http.Request is used to fetch any request specific details like json, body , or anything related to request data
func HomePage(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("this is my home page"))           // this would write to the client
	fmt.Println("home page handler function invoked") // this logs in terminal
	fmt.Println(r.URL)

}

// create Add handler function, whenever someone hits /add endpoint then do 2+2 and log in terminal and
// send the result back to the client
