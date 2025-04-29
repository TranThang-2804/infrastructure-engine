package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"your-app/internal/config"
	"your-app/internal/handler"
)

func NewRouter(cfg config.Config) *chi.Mux {
	r := chi.NewRouter()

	// Common middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Get("/health", handler.HealthCheck)

	// Example group
	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", handler.UserRoutes())
	})

	return r
}

