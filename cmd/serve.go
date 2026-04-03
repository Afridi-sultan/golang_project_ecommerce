package cmd

import (
	"eccomerce/infra/db"
	"eccomerce/repo"
	"eccomerce/rest"
	"eccomerce/rest/handlers/product"
	"eccomerce/rest/handlers/user"
	"fmt"
	"os"
)

func Serve() {
	dbcon,err := db.NewConncection()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	productRepo := repo.NewProductRepo(dbcon)
	userRepo := repo.NewUserList(dbcon)
	productHandler := product.NewHandler(productRepo)
	userHandler := user.NewHandler(userRepo)
	server := rest.NewServer(productHandler, userHandler)
	server.Start()

}
