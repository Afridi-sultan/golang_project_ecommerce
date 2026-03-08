package user

import (
	"eccomerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	//create user
	mux.Handle(
		"POST /users",
		http.HandlerFunc(h.CreateUser),
	)

	//login user
	mux.Handle(
		"POST /users/login",
		http.HandlerFunc(h.LoginUser),
	)
}
