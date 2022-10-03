package router

import (
	"Area/database"
	"Area/handlers"
	"Area/lib"
	"Area/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func ProtectedRoutes(r chi.Router) {
	tokenAuth := lib.NewTokenAuth()
	r.Use(jwtauth.Verifier(tokenAuth))
	// Todo : change this one to a custom one
	r.Use(middleware.Authenticator)
	r.Get("/triggers", handlers.Triggers)
}

func UnprotectedRoutes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		users, _ := database.User.Get(true)
		lib.SendJson(w, users)
	})
	r.Post("/login", handlers.Login)
	r.Post("/register", handlers.Register)
}
