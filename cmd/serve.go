package cmd

import (
	"eccomerce/config"
	"eccomerce/infra/db"
	"eccomerce/product"
	"eccomerce/repo"
	"eccomerce/rest"
	productHandler "eccomerce/rest/handlers/product"
	userHandler "eccomerce/rest/handlers/user"
	"eccomerce/user"
	"fmt"
	"log"
	"os"
)

func Serve() {
	// Load DB config
	dbcnf := config.DBstringLoad()
	if dbcnf == nil {
		log.Fatal("Failed to load DB config")
	}

	dbcon, err := db.NewConncection(dbcnf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Migrate the database
	errMIg := db.MigrateDB(dbcon, "./migrations")
	if errMIg != nil {
		log.Fatal("Failed to migrate database", errMIg)
		os.Exit(1)
	}

	//repository
	productRepo := repo.NewProductRepo(dbcon)
	userRepo := repo.NewUserList(dbcon)

	//services
	usrSvc := user.NewService(userRepo)
	prdctSvc := product.NewService(productRepo)

	//handlers
	productHandler := productHandler.NewHandler(prdctSvc)
	userHandler := userHandler.NewHandler(usrSvc)
	
	//server
	server := rest.NewServer(productHandler, userHandler)
	server.Start()

}
