package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leandroudala/golang_jwt/api/middleware"
)

// Route is a structure for routes
type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

// Load a route
func Load() []Route {
	routes := usersRoutes
	return routes
}

// SetupRoutes IDK what it does
func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	return r
}

// SetupRoutesWithMiddlewares adds middleware to routes
func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.URI,
			middleware.SetMiddlewareLogger(
				middleware.SetMiddlewareJSON(route.Handler),
			),
		).Methods(route.Method)
	}
	return r
}
