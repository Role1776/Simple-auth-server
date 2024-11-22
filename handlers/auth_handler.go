package handlers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/models"
	"go-rest-api/utils"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u models.Login
	_ = json.NewDecoder(r.Body).Decode(&u)
	json.NewEncoder(w).Encode(ChekLogin(u))
}

func ChekLogin(u models.Login) string {
	if models.LogUser.Username != u.Username || models.LogUser.Password != u.Password {
		fmt.Println("NOT CORRECT")
		err := "NOT CORRECT"
		return err
	}

	validToken, err := utils.GenerateJWT()
	val := "Your Token:" + validToken
	fmt.Println(validToken)

	if err != nil {
		fmt.Println(err)
	}

	return val
}
