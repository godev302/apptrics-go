package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"service-app/web"
)

func (m *Mid) Panic(next web.HandlerFunc) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
		v, ok := ctx.Value(web.KeyValue).(*web.Values)
		if !ok {
			return fmt.Errorf("web.Values missing from the context")
		}
		defer func() {
			r := recover()

			if r != nil { // panic happened
				s := fmt.Sprintf("PANIC :%v", r)
				err = errors.New(s)
				// Log the Go stack trace for this panic's goroutine.
				m.log.Printf("%s :\n%s", v.TraceId, debug.Stack())
			}

		}()

		return next(ctx, w, r)
	}
}
