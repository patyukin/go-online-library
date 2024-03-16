package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	h          *Handler
}

func New(h *Handler) *Server {
	return &Server{
		h: h,
	}
}

func (s *Server) InitRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Route("/v1/filters", func(r chi.Router) {
		r.Get("/{id}", s.h.GetFilterHandler)
		r.Delete("/{id}", s.h.GetFilterHandler)
		r.Post("/", s.h.CreatePromotionHandler)
		r.Put("/{id}", s.h.SetFilterHandler)
	})

	return r
}

func (s *Server) Run(addr string) error {
	r := s.InitRoutes()
	s.httpServer = &http.Server{
		Addr:           addr,
		Handler:        r,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
