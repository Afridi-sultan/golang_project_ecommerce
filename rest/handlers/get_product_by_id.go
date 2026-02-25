package handlers

import (
	"eccomerce/database"
	"eccomerce/util"
	"net/http"
	"strconv"
)

func GetProductByid(w http.ResponseWriter, r *http.Request){
	ProductId := r.PathValue("id")

	pID,err := strconv.Atoi(ProductId)

	if err != nil{
		http.Error(w, "Please put valid product id!",400)
		return
	}

	product := database.Get(pID)
	if product == nil{
		http.Error(w,"Product not found",404)
		return
	}
	util.SendData(w,product,200)

	
}


