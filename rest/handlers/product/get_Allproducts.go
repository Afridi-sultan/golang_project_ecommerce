package product

import (
	
	"eccomerce/util"
	"net/http"
)

// route function | product route
func (h *Handler) ProductHandler(w http.ResponseWriter, r *http.Request) {
	productList,err := h.productRepo.List()
	if err != nil{
		http.Error(w,"Internel Server Error",http.StatusBadRequest)
		return
	}
	util.SendData(w, &productList, http.StatusOK)

}
