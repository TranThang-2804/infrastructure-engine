package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", listUsers)
	r.Get("/{userID}", getUser)

	return r
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of users"))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	w.Write([]byte("User: " + userID))
}
