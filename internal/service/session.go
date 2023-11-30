package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
)

type SessionService struct {
	Secret         []byte
	ExpirationTime func() time.Time
}

func NewSessionService(secret string, expirationDuration time.Duration) *SessionService {
	return &SessionService{
		Secret: []byte(secret),
		ExpirationTime: func() time.Time {
			return time.Now().Add(expirationDuration)
		},
	}
}

func (s *SessionService) GenerateToken(email string) (string, error) {
	claims := &model.Claims{
		UserEmail: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(s.ExpirationTime()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(s.Secret)
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func (s *SessionService) ValidateToken(signedToken string) (*model.Claims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&model.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Secret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*model.Claims)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("jwt is expired")
	}

	return claims, nil
}
