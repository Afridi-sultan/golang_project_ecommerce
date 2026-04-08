package product

import (
	"eccomerce/domain"
	"eccomerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)
type ReqProducts struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}
// create product route
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct ReqProducts
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)


	if err != nil {
		fmt.Println(err)
		http.Error(w, "please give me valid json", 400)
		return
	}

	createProduct, err := h.svc.Create(domain.Products{
		Title: newProduct.Title,
		Description: newProduct.Description,
		ImgUrl: newProduct.ImgUrl,
		Price: newProduct.Price,
	})

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internel Server Error", http.StatusBadRequest)
		return
	}

	//encode new data with status code
	util.SendData(w, createProduct, 201)

}
