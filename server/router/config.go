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
	r.Get("/triggers", handlers.GetTriggers)
	r.Post("/triggers", handlers.CreateTriggers)
	r.Get("/triggers/{id}", handlers.GetTriggerById)
	r.Put("/triggers/{id}", handlers.UpdateTrigger)
	r.Delete("/triggers/{id}", handlers.DeleteTrigger)
	r.Get("/me", handlers.Me)
}

func UnprotectedRoutes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// triggers, _ := database.Trigger.Create(models.Trigger{
		// 	Title:       "test_trigger1",
		// 	Description: "Just a test trigger",
		// 	UserID:      2,
		// })
		trigger, _ := database.Trigger.Get()
		lib.SendJson(w, trigger)
	})
	r.Post("/login", handlers.Login)
	r.Post("/register", handlers.Register)
}
