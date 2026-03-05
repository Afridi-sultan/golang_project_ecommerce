package handlers

import (
	"eccomerce/database"
	"eccomerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	
	newUser.Store()
	
	//encode new data with status code
	util.SendData(w, newUser, 201)

}
