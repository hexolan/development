package http

import (
	"time"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/hexolan/stocklet/internal/app/order"
)

func NewRestAPI(svc order.OrderService) {
	r := chi.NewRouter()
	
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Add routes
	r.Get("/", getIndex)

	http.ListenAndServe(":3000", r)
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{}"))
}