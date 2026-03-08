package product

import (
	"eccomerce/database"
	"eccomerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

// create product route
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Products
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)


	if err != nil {
		fmt.Println(err)
		http.Error(w, "please give me valid json", 400)
		return
	}

	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)

	//encode new data with status code
	util.SendData(w, newProduct, 201)

}
