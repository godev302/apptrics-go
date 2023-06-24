package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	user := map[string]string{"first_name": "raj"}
	u, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}

	//NewRequest construct the request // we haven't sent it yet
	req, err := http.NewRequest(http.MethodPost, "https://httpbin.org/post", bytes.NewReader(u))

	req.Header.Set("Content-Type", "application/json") // setting headers // we will send a json to the server

	//doing the request to the remote server
	resp, err := http.DefaultClient.Do(req)
	//err is going to tell us if we are not able to call the remote service
	// it will not indicate any errors that might have happened while exec the req
	if err != nil {
		log.Fatalln(err)
	}

	//reading the response body
	bytesData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bytesData))
}

// endpoint:- https://httpbin.org/get
// body - nil
// do a get request to httpbin
