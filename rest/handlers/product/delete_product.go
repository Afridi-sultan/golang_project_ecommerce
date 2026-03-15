package product

import (
	
	"eccomerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DelProductByid(w http.ResponseWriter, r *http.Request) {
	ProductId := r.PathValue("id")

	pID, err := strconv.Atoi(ProductId)

	if err != nil {
		http.Error(w, "Please put valid product id!", 400)
		return
	}

	err2 := h.productRepo.Delete(pID)

	if err2 !=nil {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}

	productList,err := h.productRepo.List()

	if err != nil{
		http.Error(w,"internel server error",http.StatusBadRequest)
		return
	}

	util.SendData(w,&productList, 200)

}
