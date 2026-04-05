package cmd

import (
	"eccomerce/config"
	"eccomerce/infra/db"
	"eccomerce/repo"
	"eccomerce/rest"
	"eccomerce/rest/handlers/product"
	userHandler "eccomerce/rest/handlers/user"
	"fmt"
	"log"
	"os"
	"eccomerce/user"
)

func Serve() {
	dbcnf := config.DBstringLoad()
	if dbcnf == nil {
		log.Fatal("Failed to load DB config")
	}

	dbcon, err := db.NewConncection(dbcnf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	errMIg := db.MigrateDB(dbcon,"./migrations")
	if errMIg != nil {
		log.Fatal("Failed to migrate database",errMIg)
		os.Exit(1)
	}
	//repository
	productRepo := repo.NewProductRepo(dbcon)
	userRepo := repo.NewUserList(dbcon)
	//Domains
	usrSvc := user.NewService(userRepo)

	//handlers
	productHandler := product.NewHandler(productRepo)
	userHandler := userHandler.NewHandler(usrSvc)
	server := rest.NewServer(productHandler, userHandler)
	server.Start()

}
