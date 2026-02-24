package rest
import (
	"eccomerce/rest/middleware"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)
func Start() {

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
	InitRoutes(mux, manager)

	fmt.Println("Server running on :", port)
	err := http.ListenAndServe("localhost:"+port, manager.WrapMux(mux)) 

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
