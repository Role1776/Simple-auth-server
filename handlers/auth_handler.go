package handlers

import (
	"encoding/json"
	"go-rest-api/models"
	"go-rest-api/utils"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var auth models.Login

	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(checkLogin(auth))
}

func checkLogin(auth models.Login) string {
	if auth.Username != models.LogUser.Password || auth.Password != models.LogUser.Password {
		return "Invalid credentials"
	}

	validToken, err := utils.GenerateJWT()
	if err != nil {
		log.Println("Error generate token.", err)
	}
	return validToken
}
