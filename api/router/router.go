package router

import (
	"github.com/gorilla/mux"
	"github.com/leandroudala/golang_jwt/api/router/routes"
)

// New creates a new Router
func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}
