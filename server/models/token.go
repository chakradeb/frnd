package models

import "github.com/dgrijalva/jwt-go"

type Token struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Token string `json:"token"`
	*jwt.StandardClaims
}
