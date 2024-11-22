package models

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var LogUser = Login{
	Username: "1",
	Password: "1",
}