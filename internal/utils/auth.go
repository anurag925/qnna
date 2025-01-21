package utils

import "github.com/golang-jwt/jwt/v5"

var jwtKey = []byte("your-secret-key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GetJwtKey() []byte {
	return jwtKey
}
