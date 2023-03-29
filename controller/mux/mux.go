package mux

import (
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct {
	service service.Service
}

type rest struct {
	handler *handler
	server  *http.Server
	router  *mux.Router
}
