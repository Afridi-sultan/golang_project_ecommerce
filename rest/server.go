package rest

import (
	"eccomerce/rest/handlers/product"
	"eccomerce/rest/handlers/user"
	"eccomerce/rest/middleware"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Server struct{
	productHandler *product.Handler
	userHandler *user.Handler
}

func NewServer (productHandler *product.Handler, userHandler *user.Handler)*Server{
	return &Server{
		productHandler: productHandler,
		userHandler: userHandler,
	}
}

func (server *Server) Start() {

	//load env
	error := godotenv.Load()
	if error != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	//Manager
	manager := middleware.NewManager()
	manager.Use(middleware.PreflightReq, middleware.CorsMiddleWare, middleware.Logger)

	//router & routes
	mux := http.NewServeMux()
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	fmt.Println("Server running on :", port)
	err := http.ListenAndServe("localhost:"+port, manager.WrapMux(mux)) 

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
