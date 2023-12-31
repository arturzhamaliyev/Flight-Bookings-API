package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrAuthRequired = errors.New("authentication required")
	ErrAdminOnly    = errors.New("only administrator is allowed to perform this action")
	ErrCustomerOnly = errors.New("only registered customers are allowed to perform this action")

	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
)

func errorResponse(err error) any {
	return gin.H{"error": err.Error()}
}
