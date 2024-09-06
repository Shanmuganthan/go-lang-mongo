package models

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
