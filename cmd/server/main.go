package main

import (
	"net/http"
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/api/route"
	"github.com/TranThang-2804/infrastructure-engine/internal/bootstrap"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func init() {
	log.Init()

	log.Logger.Info("Starting the application - Author: Tommy Tran - tranthang.dev@gmail.com")
}

func main() {

	app := bootstrap.App()

	env := app.Env

	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

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

	route.Setup(env, timeout, r)

	log.Logger.Info("Starting server...", "on port", env.ServerAddress)
	http.ListenAndServe(env.ServerAddress, r)
}

func chatRouter() http.Handler {
	r := chi.NewRouter()
	// r.Post("/", handler.blueprint)
	return r
}
