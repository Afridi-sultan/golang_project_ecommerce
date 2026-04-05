package user

import (
	"eccomerce/config"
	"eccomerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string
	Password string
}

// create product route
func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {

	var loginUser ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginUser)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	usr, err := h.svc.GetUser(loginUser.Email,loginUser.Password)

	jwtSecKey := config.SecretKey

	accessToken, e := util.CreateJwt(jwtSecKey, util.Payloader{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})

	if e != nil {
		http.Error(w, "internel server error", 401)
		return
	}

	util.SendData(w, accessToken, 200)

}
