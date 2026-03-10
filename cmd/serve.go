package cmd

import (
	"eccomerce/rest"
	"eccomerce/rest/handlers/product"
	"eccomerce/rest/handlers/user"

)

func Serve() {
	productHandler := product.NewHandler()
	userHandler := user.NewHandler()
	server := rest.NewServer(productHandler, userHandler)
	server.Start()

}
