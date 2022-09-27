package router

import (
	"Area/handlers"
	"Area/lib"
	"errors"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/oauth"
)

func ProtectedRoutes(r chi.Router) {
	r.Use(oauth.Authorize(os.Getenv("TOKEN_SECRET"), nil))
	r.Get("/triggers", handlers.Triggers)
}

func UnprotectedRoutes(r chi.Router) {
	r.Get("/login", handlers.Login)
	r.Get("/register", handlers.Register)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		lib.CheckError(errors.New("Agneugneugneu"))
	})
}
