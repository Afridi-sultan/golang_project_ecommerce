package handlers

import (
	"eccomerce/database"
	"eccomerce/util"
	"net/http"
)

// route function | product route
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.ProductList, 200)

}
