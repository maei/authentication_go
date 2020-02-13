package models

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct{
	jwt.StandardClaims
}

func CreateToken(uuid string) (string, string) {
	return "refresh_token", "access_token"
}