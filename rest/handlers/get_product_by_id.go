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
	util.SendData(w,product,200)

	// for _,product := range database.ProductList{
	// 	if product.ID == pID{
	// 		util.SendData(w,product,200)
	// 	}
	// }
}