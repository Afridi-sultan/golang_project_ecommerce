package cmd

import (
	"eccomerce/repo"
	"eccomerce/rest"
	"eccomerce/rest/handlers/product"
	"eccomerce/rest/handlers/user"
)

func Serve() {
	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserList()
	productHandler := product.NewHandler(productRepo)
	userHandler := user.NewHandler(userRepo)
	server := rest.NewServer(productHandler, userHandler)
	server.Start()

}
