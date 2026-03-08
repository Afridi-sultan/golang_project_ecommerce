package product

import (
	"eccomerce/database"
	"eccomerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProductByid(w http.ResponseWriter, r *http.Request) {
	ProductId := r.PathValue("id")

	pID, err := strconv.Atoi(ProductId)

	if err != nil {
		http.Error(w, "Please put valid product id!", 400)
		return
	}

	var updateProduct database.Products

	decoder := json.NewDecoder(r.Body)

	errJson := decoder.Decode(&updateProduct)
	if errJson != nil {
		http.Error(w, "Please put valid product id!", 400)
		return
	}

	updated := database.Update(pID, updateProduct)
	if updated == nil {
		http.Error(w, "product not found", 404)
	}
	util.SendData(w, updateProduct, 200)

}
