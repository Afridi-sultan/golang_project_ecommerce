package rest

import (
	"eccomerce/rest/handlers"
	"eccomerce/rest/middleware"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /products",
		manager.With(middleware.Hudai)(
			http.HandlerFunc(handlers.ProductHandler),
		),
	)

	mux.Handle(
		"POST /create-products",
		manager.With(middleware.AuthenticateJwt)(
			http.HandlerFunc(handlers.CreateProduct),
		),
	)

	mux.Handle(
		"PUT /products/{id}", 
		manager.With(middleware.AuthenticateJwt)(
			http.HandlerFunc(handlers.UpdateProductByid),
		),
	)

	mux.Handle("DELETE /products/{id}", 
	manager.With(middleware.AuthenticateJwt)(
		http.HandlerFunc(handlers.DelProductByid),
	),
	)

	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductByid))
	mux.Handle("POST /users", http.HandlerFunc(handlers.CreateUser))
	mux.Handle("POST /users/login", http.HandlerFunc(handlers.LoginUser))
}
