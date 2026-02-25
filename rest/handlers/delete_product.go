package handlers

import (
	"eccomerce/database"
	"eccomerce/util"
	"net/http"
	"strconv"
)

func DelProductByid(w http.ResponseWriter, r *http.Request) {
	ProductId := r.PathValue("id")

	pID, err := strconv.Atoi(ProductId)

	if err != nil {
		http.Error(w, "Please put valid product id!", 400)
		return
	}

	delete := database.Delete(pID)

	if !delete {
		http.Error(w, "page not found", 404)
		return
	}

	util.SendData(w, database.GetAllProduct(), 200)

}
