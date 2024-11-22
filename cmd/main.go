package main

import (
	"go-rest-api/handlers"
	"go-rest-api/middleware"
	"go-rest-api/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	models.Users = append(models.Users, models.User{
		Id:       "1",
		Name:     "Michael Johnson",
		Age:      20,
		NameBook: "The Go Programming Guide",
		Book:     models.Book{Theme: "Introduction to Go", Year: 2023},
	})

	models.Users = append(models.Users, models.User{
		Id:       "2",
		Name:     "John Smith",
		Age:      19,
		NameBook: "World History: An Overview",
		Book:     models.Book{Theme: "History of Civilizations", Year: 1999},
	})

	router := mux.NewRouter()

	router.Handle("/usersBook", middleware.ChekAuth(handlers.GetUsers)).Methods("GET")
	router.Handle("/usersBook/{id}", middleware.ChekAuth(handlers.GetUsersById)).Methods("GET")
	router.Handle("/usersBook", middleware.ChekAuth(handlers.CreateUser)).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.Handle("/usersBook/{id}", middleware.ChekAuth(handlers.PutUsers)).Methods("PUT")
	router.Handle("/usersBook/{id}", middleware.ChekAuth(handlers.DeleteUsers)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}