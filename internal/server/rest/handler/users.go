package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest/handler/request"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest/handler/response"
)

// Users represents a type that provides operations on users.
type Users interface {
	CreateUser(ctx context.Context, user model.User) error
}

// CreateUserss will try to create user, responses with Created status and Created user info if no error occured.
func (h *Handler) CreateUser(ctx *gin.Context) {
	var userReq request.CreateUser
	err := ctx.ShouldBindJSON(&userReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		ID:        uuid.New(),
		Phone:     userReq.Phone,
		Email:     userReq.Email,
		Password:  userReq.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = h.usersService.CreateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := response.CreateUser{
		ID:    user.ID,
		Phone: user.Phone,
		Email: user.Email,
	}

	ctx.JSON(http.StatusCreated, resp)
}
