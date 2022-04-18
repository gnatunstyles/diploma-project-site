package handlers

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	EncryptedPwd string `json:"pwd"`
}
