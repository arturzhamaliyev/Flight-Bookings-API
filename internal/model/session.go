package model

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserEmail string `json:"userEmail"`
	jwt.RegisteredClaims
}
