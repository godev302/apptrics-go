package web

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type App struct {
	*chi.Mux // Mux impls http.Handler interface
}

// HandlerFunc is a custom type like http.HandlerFunc func in standard library
type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func (a *App) HandleFunc(method string, pattern string, handler HandlerFunc) {
	f := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context() // taking context out of the Request
		err := handler(ctx, w, r)
		if err != nil {
			log.Println("error escaped from the middleware ", err)
			return
		}
	}

	// chi router can accept the f var because the signature of f var matches
	//to func(w http.ResponseWriter, r *http.Request)
	a.Mux.MethodFunc(method, pattern, f)
}
