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
)

// Users represents a type that provides operations on users.
type Users interface {
	CreateUser(ctx context.Context, user model.User) error
}

// CreateUserss will try to create user, responses with Created status and Created user info if no error occured.
// Will replace with Swagger soon.
func (h *Handler) CreateUser(ctx *gin.Context) {
	var userReq request.CreateUser
	err := ctx.ShouldBindJSON(&userReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
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
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, response.CreateUser{
		ID:    user.ID,
		Phone: user.Phone,
		Email: user.Email,
	})
}
