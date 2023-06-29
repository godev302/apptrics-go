package handlers

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"service-app/auth"
	"service-app/middleware"
	"service-app/web"
)

//go get github.com/go-chi/chi/v5

func Api(log *log.Logger, a *auth.Auth) http.Handler {
	//r := chi.NewRouter()

	app := web.App{
		Mux: chi.NewRouter(),
	}
	m := middleware.NewMid(log, a)
	app.HandleFunc(http.MethodGet, "/check", m.Logger(m.Error(m.Panic(m.Authenticate(m.Authorize(check, auth.RoleAdmin))))))

	return app
}
