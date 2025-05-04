package main

import (
	"net/http"
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/api/route"
	"github.com/TranThang-2804/infrastructure-engine/internal/bootstrap"
	"github.com/TranThang-2804/infrastructure-engine/internal/mq"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/env"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func init() {
	log.Init()
	env.LoadEnv()

	log.Logger.Info("Starting the application - Author: Tommy Tran - tranthang.dev@gmail.com")
}

func main() {

	app := bootstrap.App()

	gitStore := app.GitStore

	defer app.CloseDBConnection()

	timeout := time.Duration(env.Env.ContextTimeout) * time.Second

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

	route.Setup(gitStore, timeout, r)
  mq.SetupMQController(gitStore, timeout)

	log.Logger.Info("Starting server...", "on port", env.Env.ServerAddress)
	http.ListenAndServe(env.Env.ServerAddress, r)
}
