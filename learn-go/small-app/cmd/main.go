package main

import (
	"net/http"
	"small-app/handlers"
)

// http://localhost:8080/search?user_id=123
func main() {
	http.HandleFunc("/search", handlers.GetUser)
	panic(http.ListenAndServe(":8080", nil))
}
