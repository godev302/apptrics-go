package handlers

import (
	"context"
	"net/http"
	"service-app/web"
)

func check(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string // field
	}{
		Status: "ok",
	}
	_ = status
	//panic("some panic")
	//ErrSomething := errors.New("some trusted error")
	return web.Respond(ctx, w, status, http.StatusOK)

	//return web.NewRequestError(ErrSomething, http.StatusBadRequest)

	//return errors.New("some error")
	//
	//return json.NewEncoder(w).Encode(status)

}
