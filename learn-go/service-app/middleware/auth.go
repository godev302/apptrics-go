package middleware

import (
	"context"
	"errors"
	"net/http"
	"service-app/auth"
	"service-app/web"
	"strings"
)

var ErrForbidden = web.NewRequestError(
	errors.New("you are not authorized for that action"),
	http.StatusForbidden,
)

func (m *Mid) Authenticate(next web.HandlerFunc) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		// fetching data from Authorization header
		authHeader := r.Header.Get("Authorization")
		//token format :- bearer <token>

		parts := strings.Split(authHeader, " ") // parts would be slice
		claims, err := m.a.ValidateToken(parts[1])
		if err != nil {
			return web.NewRequestError(err, http.StatusUnauthorized)
		}
		ctx = context.WithValue(ctx, auth.Key, claims)
		return next(ctx, w, r)
	}
}

func (m *Mid) Authorize(next web.HandlerFunc, requiredRoles ...string) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		claims, ok := ctx.Value(auth.Key).(auth.Claims)
		if !ok {
			return errors.New("claims missing from context: Authorize called without/before Authenticate")
		}
		ok = claims.HasRoles(requiredRoles...)
		if !ok {
			return ErrForbidden
		}
		return next(ctx, w, r)
	}
}
