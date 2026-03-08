package product

import (
	
	"eccomerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /products",
		manager.With(middleware.Hudai)(
			http.HandlerFunc(h.ProductHandler),
		),
	)

	mux.Handle(
		"POST /create-products",
		manager.With(middleware.AuthenticateJwt)(
			http.HandlerFunc(h.CreateProduct),
		),
	)

	mux.Handle(
		"PUT /products/{id}",
		manager.With(middleware.AuthenticateJwt)(
			http.HandlerFunc(h.UpdateProductByid),
		),
	)

	mux.Handle(
		"DELETE /products/{id}",
		manager.With(middleware.AuthenticateJwt)(
			http.HandlerFunc(h.DelProductByid),
		),
	)

	mux.Handle("GET /products/{id}", http.HandlerFunc(h.GetProductByid))

}
