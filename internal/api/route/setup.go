package route

import (
	"github.com/TranThang-2804/infrastructure-engine/internal/api/middleware"
	"github.com/TranThang-2804/infrastructure-engine/internal/bootstrap"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func SetupRoute(app bootstrap.Application) *chi.Mux {
	r := chi.NewRouter()

	// CORS middleware
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"}, // Use your allowed origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Origin", "X-Requested-With"},
		ExposedHeaders:   []string{"Link", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	// Define middleware
	r.Use(middleware.RealIP)
	r.Use(custommiddleware.LoggingMiddleware)
	r.Use(middleware.Recoverer)

	// Public APIs
	r.Group(func(r chi.Router) {
		NewHealthRouter(r, app.HealthController)
		NewBluePrintRouter(r, app.BluePrintController)
		NewIacTemplateRouter(r, app.IacTemplateController)
		NewCompositeResourceRouter(r, app.CompositeResourceController)
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		// r.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	})

	return r
}
