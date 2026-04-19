package product

import (
	"eccomerce/domain"
	"eccomerce/util"
	"net/http"
	"strconv"

	"sync"
)

var countMain int64
var mu sync.Mutex

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

	productCh := make(chan []*domain.Products, 1) // Create a buffered channel with capacity 1
	go func() {
		productList, err := h.svc.List(pageNum, limitNum)
		productCh <- productList // Send the productList to the channel
			if err != nil {
			http.Error(w, "Internel Server Error", http.StatusBadRequest)
			return
		}

	}()

	ch := make(chan int64, 1) // Create a buffered channel with capacity 1

	go func() {
		count, err := h.svc.Count()
		ch <- count // Send the count to the channel
		if err != nil {
			http.Error(w, "Internel Server Error", http.StatusBadRequest)
			return
		}

	}()

	countMain := <-ch // Receive the count from the channel
	productList := <-ch // Receive the productList from the channel

	util.SendPaginatedData(w, &productList, pageNum, limitNum, countMain)

}
