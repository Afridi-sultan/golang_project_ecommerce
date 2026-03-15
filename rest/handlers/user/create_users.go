package user

import (
	
	"eccomerce/repo"
	"eccomerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqUser struct{
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsShopOwner bool `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser ReqUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}


	usr, err := h.userRepo.Create(&repo.User{
		FirstName: newUser.FirstName,
		Email: newUser.Email,
		Password: newUser.Password,
	})
	if err != nil {
		http.Error(w, "Internel server error", http.StatusBadRequest)
		return
	}
	
	//encode new data with status code
	util.SendData(w, &usr, http.StatusCreated)

}
