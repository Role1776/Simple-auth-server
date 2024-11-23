package utils

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var MySignKey = []byte("oyminirole")


func GenerateJWT() (string,error) {
	token := jwt.New(jwt.SigningMethodHS384)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 800).Unix()
	claims["user"] = "oyminirole"

	stringToken,err := token.SignedString(MySignKey)
	if err != nil {
		log.Println(err)
	}
	
	return stringToken, nil
}
