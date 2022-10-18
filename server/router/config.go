package router

import (
	"Area/handlers"
	"Area/lib"
	"Area/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func ProtectedRoutes(r chi.Router) {
	tokenAuth := lib.NewTokenAuth()
	r.Use(jwtauth.Verifier(tokenAuth))
	// Todo : change this one to a custom one
	r.Use(middleware.Authenticator)
	r.Get("/triggers", handlers.GetTriggers)
	r.Post("/triggers", handlers.CreateTriggers)
	r.Get("/triggers/{id}", handlers.GetTriggerById)
	r.Put("/triggers/{id}", handlers.UpdateTrigger)
	r.Delete("/triggers/{id}", handlers.DeleteTrigger)
	r.Get("/me", handlers.Me)
	r.Get("/providers/{provider}/auth", handlers.ProviderLogin)
	r.Get("/actions", handlers.GetActions)
	r.Get("/reactions", handlers.GetReactions)
}

func UnprotectedRoutes(r chi.Router) {
	r.Post("/login", handlers.Login)
	r.Post("/register", handlers.Register)
	r.Get("/providers/{provider}/callback", handlers.ProviderCallback)
	r.Get("/logout", handlers.Logout)
}
