package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

var r *chi.Mux

func HandleFunc(pattern string, handler HandlerFunc) {
	f := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := handler(ctx, w, r)
		if err != nil {
			log.Fatal(err)
		}
	}

	r.HandleFunc(pattern, f)

}

func main() {

	r = chi.NewRouter()
	HandleFunc("/home", Home)
	http.ListenAndServe(":8081", r)
	//r.HandleFunc("/home", Home)
}

func Home(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintln(w, "testing custom handlers")
	return nil
}
