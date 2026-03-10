package rest

import (
	"eccomerce/config"
	"eccomerce/rest/handlers/product"
	"eccomerce/rest/handlers/user"
	"eccomerce/rest/middleware"
	"fmt"

	"net/http"
)

type Server struct {
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
	config.LoadConfig()

	//Manager
	manager := middleware.NewManager()
	manager.Use(middleware.PreflightReq, middleware.CorsMiddleWare, middleware.Logger)

	//router & routes
	mux := http.NewServeMux()
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	fmt.Println("Server running on :", config.Port)
	err := http.ListenAndServe("localhost:"+config.Port, manager.WrapMux(mux))

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
