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
	http.Error(w, "User not found", http.StatusNotFound)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var us models.User
	err := json.NewDecoder(r.Body).Decode(&us)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	models.Users = append(models.Users, us)
	json.NewEncoder(w).Encode(models.Users)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range models.Users {
		if item.Id == params["id"] {
			models.Users = append(models.Users[:i], models.Users[i+1:]...)
			json.NewEncoder(w).Encode(models.Users)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func PutUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	var us models.User

	err := json.NewDecoder(r.Body).Decode(&us)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	for i, item := range models.Users {
		if item.Id == params["id"] {
			us.Id = params["id"]
			models.Users[i] = us
			json.NewEncoder(w).Encode(models.Users)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}


