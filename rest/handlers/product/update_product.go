package product

import (
	"eccomerce/repo"
	"eccomerce/util"
	"encoding/json"
	"net/http"
	"strconv"
	// "strconv"
)
type Rproducts struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}
func (h *Handler) UpdateProductByid(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pId, err := strconv.Atoi(productId)

	var updateProduct Rproducts

	decoder := json.NewDecoder(r.Body)

	errJson := decoder.Decode(&updateProduct)
	if errJson != nil {
		http.Error(w, "Please put valid product id!", 400)
		return
	}

	updated, err := h.productRepo.Update(repo.Products{
		ID: pId,
		Title:updateProduct.Title ,
		Description: updateProduct.Description,
		ImgUrl: updateProduct.ImgUrl,
		Price: updateProduct.Price,
	})

	if err !=nil{
		http.Error(w,"Internel server error",http.StatusBadRequest)
	}
	
	util.SendData(w, &updated, http.StatusOK)

}
