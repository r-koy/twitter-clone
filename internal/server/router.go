package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"twitter-clone/api/v1/health"
	"twitter-clone/api/v1/user"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", health.Ping)
	r.Get("/users/{id}", user.GetUserById)

	return r
}
