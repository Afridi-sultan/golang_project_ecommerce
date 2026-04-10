package product

import (
	
	"eccomerce/util"
	"net/http"
	"strconv"
)



// route function | product route
func (h *Handler) ProductHandler(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	page := queryParam.Get("page")
	pageNum, _ := strconv.ParseInt(page, 10, 32)
	limit := queryParam.Get("limit")
	limitNum, _ := strconv.ParseInt(limit, 10, 32)

	if pageNum == 0 {
		pageNum = 1
	}
	if limitNum == 0 {
		limitNum = 10
	}

	productList, err := h.svc.List(pageNum, limitNum)
	if err != nil {
		http.Error(w, "Internel Server Error", http.StatusBadRequest)
		return
	}

	count, err := h.svc.Count()
	if err != nil {
		http.Error(w, "Internel Server Error", http.StatusBadRequest)
		return
	}



		util.SendPaginatedData(w,&productList,pageNum,limitNum,count)

}
