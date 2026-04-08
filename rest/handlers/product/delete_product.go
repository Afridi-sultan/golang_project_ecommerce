package product

import (
	
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

	err2 := h.svc.Delete(pID)

	if err2 !=nil {
		if err2.Error() == "product not found"{
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
