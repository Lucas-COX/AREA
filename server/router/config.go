package router

import (
	"Area/handlers"
	"Area/lib"
	"Area/middleware"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/oauth"
)

func ProtectedRoutes(r chi.Router) {
	r.Use(oauth.Authorize(os.Getenv("TOKEN_SECRET"), nil))
	r.Get("/triggers", handlers.Triggers)
}

func UnprotectedRoutes(r chi.Router) {
	s := oauth.NewBearerServer(
		os.Getenv("TOKEN_SECRET"),
		time.Second*600,
		&middleware.UserVerifier{},
		nil,
	)

	r.Use(middleware.JsonToForm)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		lib.CheckError(errors.New("Agneugneugneu"))
	})
	r.Post("/register", handlers.Register)
	r.Post("/token", s.UserCredentials)
	r.Post("/auth", s.ClientCredentials)
}
