package models

type Book struct {
	Theme string `json:"name"`
	Year  int    `json:"year"`
}

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	NameBook string `json:"name_book"`
	Book     Book   `json:"book"`
}

var Users []User