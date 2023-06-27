package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"service-app/web"
)

//go get github.com/go-chi/chi/v5

func Api() http.Handler {
	//r := chi.NewRouter()

	app := web.App{
		Mux: chi.NewRouter(),
	}

	app.HandleFunc(http.MethodGet, "/check", check)

	return app
}
