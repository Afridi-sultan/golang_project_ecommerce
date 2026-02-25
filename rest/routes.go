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

	mux.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductByid))
	mux.Handle("PUT /products/{id}", http.HandlerFunc(handlers.UpdateProductByid))
	mux.Handle("DELETE /products/{id}", http.HandlerFunc(handlers.DelProductByid))
}
