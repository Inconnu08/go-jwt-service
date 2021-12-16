package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"socialmedia-auth/internal/service"
)

type handler struct {
	*service.Service
}

func New(s *service.Service) http.Handler {
	h := &handler{s}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/", func(r chi.Router) {
		r.Get("/auth/readme", h.getReadme)
		r.Get("/auth/README.txt", h.getReadme)
		r.Get("/auth/stats", h.getStats)
		r.Get("/auth/{username}", h.createToken)
		r.Get("/verify", h.verifyToken)
		r.Get("/readme", h.getReadme)
		r.Get("/stats", h.getStats)
		r.Get("/README.txt", h.getReadme)
	})
	return r
}
