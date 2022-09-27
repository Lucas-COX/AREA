package middleware

import (
	"Area/handlers"
	"Area/lib"
	"net/http"
	"runtime/debug"

	"github.com/go-chi/chi/v5/middleware"
)

func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {

				logEntry := middleware.GetLogEntry(r)
				if logEntry != nil {
					logEntry.Panic(rvr, debug.Stack())
				} else {
					middleware.PrintPrettyStack(rvr)
				}

				w.WriteHeader(http.StatusInternalServerError)
				lib.SendJson(w, handlers.ErrorBody{
					Message: "An unexpected error occurred",
				})
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
