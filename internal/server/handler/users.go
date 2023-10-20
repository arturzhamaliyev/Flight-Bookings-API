package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model/request"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Users represents a type that provides operations on users.
type Users interface {
	CreateUser(ctx context.Context, user model.User) error
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	var userReq request.CreateUser
	err := ctx.ShouldBindJSON(&userReq)
	if err != nil {
		zap.S().Infof("failed to bind json: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	user := model.User{
		ID:        uuid.New(),
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Password:  userReq.Password,
		Email:     userReq.Email,
		Country:   userReq.Country,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = h.usersService.CreateUser(ctx, user)
	if err != nil {
		zap.S().Infof("failed to create user: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, response.CreateUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Country:   user.Email,
	})
}
