package util

import "net/http"

type PaginatedData struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalItems int64 `json:"total_items"`
	TotalPages int64 `json:"total_pages"`
}

func SendPaginatedData(w http.ResponseWriter, data any, pageNum, limitNum, count int64) {
		paginatedData := PaginatedData{

		Data:       data,

		Pagination: Pagination{
			Page:       pageNum,
			Limit:      limitNum,
			TotalItems: count, 
			TotalPages: count / limitNum, 
		},
		
	}

	SendData(w,&paginatedData,http.StatusOK)
}