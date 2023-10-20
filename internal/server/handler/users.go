package handler

import (
	"context"
	"net/http"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Users represents a type that provides operations on users.
type Users interface {
	CreateUser(ctx context.Context, user model.CreateUserRequest) error
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	var userReq model.CreateUserRequest
	err := ctx.ShouldBindJSON(&userReq)
	if err != nil {
		zap.S().Infof("failed to bind json: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = h.users.CreateUser(ctx, userReq)
	if err != nil {
		zap.S().Infof("failed to create user: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, "")
}
