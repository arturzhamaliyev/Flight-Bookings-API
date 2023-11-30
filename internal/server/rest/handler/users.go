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

// UsersService represents a type that provides operations on users.
type UsersService interface {
	CreateUser(ctx context.Context, user model.User) error
	CheckUserCredentials(ctx context.Context, email, password string) error
}

type SessionService interface {
	GenerateToken(email string) (string, error)
	ValidateToken(signedToken string) (*model.Claims, error)
}

// SignUp will try to create user, responses with Created status and Created user info if no error occured.
func (h *Handler) SignUp(ctx *gin.Context) {
	var userReq request.SignUp
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

	resp := response.SignUp{
		ID:    user.ID,
		Phone: user.Phone,
		Email: user.Email,
	}

	ctx.JSON(http.StatusCreated, resp)
}

// SignIn
func (h *Handler) SignIn(ctx *gin.Context) {
	var userReq request.SignIn
	err := ctx.ShouldBindJSON(&userReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.usersService.CheckUserCredentials(ctx, userReq.Email, userReq.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := h.sessionService.GenerateToken(userReq.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := response.SignIn{
		Token: token,
	}

	ctx.JSON(http.StatusOK, resp)
}
