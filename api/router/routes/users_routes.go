package routes

import (
	"net/http"

	"github.com/leandroudala/golang_jwt/api/controllers/users"
)

var usersRoutes = []Route{
	{
		URI:     "/user",
		Method:  http.MethodGet,
		Handler: users.All,
	},
	{
		URI:     "/user",
		Method:  http.MethodPost,
		Handler: users.Create,
	},
	{
		URI:     "/user/{id}",
		Method:  http.MethodPut,
		Handler: users.Update,
	},
	{
		URI:     "/user/{id}",
		Method:  http.MethodDelete,
		Handler: users.Delete,
	},
	{
		URI:     "/user/{id}",
		Method:  http.MethodGet,
		Handler: users.Get,
	},
}
