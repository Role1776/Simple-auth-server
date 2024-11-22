package handlers

import (
	"encoding/json"
	"go-rest-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(models.Users)
}

func GetUsersById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range models.Users {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "json/application")
	var us models.User
	_ = json.NewDecoder(r.Body).Decode(&us)
	models.Users = append(models.Users, us)
	json.NewEncoder(w).Encode(models.Users)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "json/application")
	params := mux.Vars(r)
	for i, item := range models.Users {
		if item.Id == params["id"] {
			models.Users = append(models.Users[:i], models.Users[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(models.Users)
}

func PutUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "json/application")
	params := mux.Vars(r)
	for i, item := range models.Users {
		if item.Id == params["id"] {
			var us models.User
			models.Users = append(models.Users[:i], models.Users[i+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&us)
			us.Id = params["id"]
			models.Users = append(models.Users, us)
			json.NewEncoder(w).Encode(models.Users)
		}
	}
}