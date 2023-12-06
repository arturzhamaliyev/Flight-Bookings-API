package helper

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	customErrors "github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/errors"
)

func GetCurrentUserIDFromToken(ctx *gin.Context) (uuid.UUID, error) {
	token, err := GetToken(ctx)
	if err != nil {
		return uuid.UUID{}, err
	}

	err = ValidateToken(token)
	if err != nil {
		return uuid.UUID{}, err
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)

	return uuid.Parse(userID)
}

func GenerateToken(user model.User) (string, error) {
	var role model.Role

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":          user.ID,
			role.String(): user.Role,
			"iat":         time.Now().Unix(),
			"eat":         time.Now().Add(model.ExpirationDuration),
		},
	)
	return token.SignedString([]byte(model.SecretKey))
}

func ValidateToken(token *jwt.Token) error {
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return customErrors.ErrInvalidToken
}

func ValidateAdminRole(token *jwt.Token) error {
	var role model.Role
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := model.Role(claims[role.String()].(float64))
	if ok && token.Valid && userRole == model.Admin {
		return nil
	}
	return customErrors.ErrNotAdminRole
}

func ValidateCustomerRole(token *jwt.Token) error {
	var role model.Role
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := model.Role(claims[role.String()].(float64))
	if ok && token.Valid && userRole == model.Customer {
		return nil
	}
	return customErrors.ErrNotCustomer
}

func GetToken(ctx *gin.Context) (*jwt.Token, error) {
	return jwt.Parse(ExtractTokenFromRequest(ctx), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(model.SecretKey), nil
	})
}

func ExtractTokenFromRequest(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	splitedToken := strings.Split(bearerToken, " ")
	if len(splitedToken) == 2 {
		return splitedToken[1]
	}
	return ""
}
