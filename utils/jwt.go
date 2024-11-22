package utils

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var MySignKey = []byte("oyminirole")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 1000).Unix()
	claims["user"] = "oyminirole"
	claims["authorized"] = true

	tokenString, err := token.SignedString(MySignKey)

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}