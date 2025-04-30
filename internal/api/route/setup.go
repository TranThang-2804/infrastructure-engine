package route

import (
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/adapter/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/bootstrap"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Setup(env *bootstrap.Env, gitStore git.GitStore, timeout time.Duration, r *chi.Mux) {
	// Define middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Public APIs
	r.Group(func(r chi.Router) {
		NewHealthCheckRouter(env, timeout, r)
		NewBluePrintRouter(env, gitStore, timeout, r)
    NewIacTemplateRouter(env, gitStore, timeout, r)
		// NewSignupRouter(env, timeout, db, r)
		// NewLoginRouter(env, timeout, db, r)
		// NewRefreshTokenRouter(env, timeout, db, r)
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		// r.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
		// NewTaskRouter(env, timeout, db, r)
		// NewProfileRouter(env, timeout, db, r)
	})
}
