package product

import (
	"eccomerce/util"
	"net/http"
	"strconv"

	"sync"
)

var countMain int64

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

	var wg sync.WaitGroup

	//just for testing purpose
	wg.Add(1)
	go func() {
		defer wg.Done()
		count, err := h.svc.Count()
		countMain = count
		if err != nil {
			http.Error(w, "Internel Server Error", http.StatusBadRequest)
			return
		}
		
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		count1, err := h.svc.Count()
		countMain = count1
		if err != nil {
			http.Error(w, "Internel Server Error", http.StatusBadRequest)
			return
		}
		
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		count2, err := h.svc.Count()
		countMain = count2

		if err != nil {
			http.Error(w, "Internel Server Error", http.StatusBadRequest)
			return
		}
		
	}()
	
	wg.Wait()
	//just for testing purpose

	util.SendPaginatedData(w, &productList, pageNum, limitNum, countMain)

}
