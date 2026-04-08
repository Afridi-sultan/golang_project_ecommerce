package product

import (
	
	"eccomerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductByid(w http.ResponseWriter, r *http.Request){
	ProductId := r.PathValue("id")

	pID,err := strconv.Atoi(ProductId)

	if err != nil{
		http.Error(w, "Please put valid product id!",400)
		return
	}

	
	product, err := h.svc.Get(pID)
	if err != nil{
		http.Error(w,"Internal server error",http.StatusBadRequest)
		return
	}
	
	util.SendData(w,&product,200)
	

	
}


