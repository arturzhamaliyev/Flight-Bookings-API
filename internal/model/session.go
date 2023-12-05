package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	ExpirationDuration = 10 * time.Minute
	SecretKey          = "secret"
)

type Claims struct {
	UserEmail string `json:"userEmail"`
	jwt.RegisteredClaims
}
