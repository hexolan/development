package rest

import (
	"github.com/gorilla/mux"
)

func NewRESTServer() {
	r := mux.NewRouter()
	r.HandleFunc("/todo", TodoRouteHandler)
}
