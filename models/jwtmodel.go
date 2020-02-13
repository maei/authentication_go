package models

import (
	"time"
	"log"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	Uuid string `json:"userId"`
	jwt.StandardClaims
}

func CreateToken(uuid string) (string, error) {
	claims := JWTClaims{
		uuid,
		jwt.StandardClaims{
			Id: "main_userId",
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Hash the token
	// add .env file-path
	token, err := rawToken.SignedString([]byte("superSecret"))
	if err != nil{
		log.Print(err)
		return "", err
	}

	return token, nil
}
