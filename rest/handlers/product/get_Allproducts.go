package product

import (
	"eccomerce/database"
	"eccomerce/util"
	"net/http"
)

// route function | product route
func (h *Handler) ProductHandler(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.ProductList, 200)

}
