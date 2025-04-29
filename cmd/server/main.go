package main

import (
	"github.com/TranThang-2804/infrastructure-engine/shared/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func init() {
	log.Init()

	log.Logger.Info("Starting the application - Author: Tommy Tran - tranthang.dev@gmail.com")
}

func main() {
	r := chi.NewRouter()

	// Define middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/api", chatRouter())

	http.ListenAndServe(":3000", r)

	log.Logger.Info("Starting server on :8080...")
}

func chatRouter() http.Handler {
	r := chi.NewRouter()
	// r.Post("/", handler.blueprint)
	return r
}
