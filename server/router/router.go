package router

import (
	"Area/handlers"
	m "Area/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(m.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(m.AutoHeaders(map[string]string{
		"Content-Type": "application/json",
	}))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(middleware.Heartbeat("/ping"))

	r.Group(ProtectedRoutes)
	r.Group(UnprotectedRoutes)

	r.NotFound(handlers.NotFound)
	return r
}
