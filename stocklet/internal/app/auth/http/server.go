package http

import (
	"time"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/hexolan/stocklet/internal/app/auth"
)

func NewRestAPI(svc auth.AuthService) {
	r := chi.NewRouter()
	
	// Add middleware
	// r.Use(middleware.RequestID) (todo: request ids at gateway - prometheus tracing, etc)
	// r.Use(middleware.Logger) (todo: use zerolog)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))
	// todo: ensuring content response headers, etc...

	// Add routes
	r.Get("/", getIndex)

	http.ListenAndServe(":3000", r)
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{}"))
}