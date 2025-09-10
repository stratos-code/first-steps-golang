package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"

	"goapi/internal/middleware"
)

// The name has a capital H because it's used in cmd/api/main.go
func Handler(r *chi.Mux) {
	// Global middlewares
	r.Use(chimiddle.StripSlashes) // Strip slashes from request URL (e.g. /foo/ -> /foo)

	r.Route("/account", func(router chi.Router) {

		// Middlewares for /account routes
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
