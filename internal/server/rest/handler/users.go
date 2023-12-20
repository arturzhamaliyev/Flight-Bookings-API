package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/helper"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest/handler/request"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest/handler/response"
)

// UsersService represents a type that provides operations on users.
type UsersService interface {
	GetUserByID(ctx context.Context, ID uuid.UUID) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	CreateUser(ctx context.Context, user model.User) error
	ValidateUserPassword(hashedPassword, password string) error
	UpdateUser(ctx context.Context, user model.User) error
	DeleteUserByID(ctx context.Context, ID uuid.UUID) error
}

// SignUp will try to create user, responses with Created status and Created user info if no error occured.
func (h *Handler) SignUp(ctx *gin.Context) {
	var userReq request.SignUp
	err := ctx.ShouldBindJSON(&userReq)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user := model.User{
		ID:        uuid.New(),
		Role:      model.Customer,
		Phone:     userReq.Phone,
		Email:     userReq.Email,
		Password:  userReq.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = h.usersService.CreateUser(ctx, user)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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
		zap.S().Info(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := h.usersService.GetUserByEmail(ctx, userReq.Email)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusUnauthorized, errorResponse(ErrInvalidEmailOrPassword))
		return
	}

	err = h.usersService.ValidateUserPassword(user.Password, userReq.Password)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusUnauthorized, errorResponse(ErrInvalidEmailOrPassword))
		return
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(model.ExpirationDuration),
	})
}

// SignOut
func (h *Handler) SignOut(ctx *gin.Context) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:   "token",
		MaxAge: -1,
	})
}

func (h *Handler) UpdateProfile(ctx *gin.Context) {
	userID, err := helper.GetCurrentUserIDFromToken(ctx)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	var userReq request.UpdateProfile
	err = ctx.ShouldBindJSON(&userReq)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user := model.User{
		ID:       userID,
		Phone:    userReq.Phone,
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	err = h.usersService.UpdateUser(ctx, user)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := response.UpdateProfile{
		ID:        user.ID,
		Phone:     user.Phone,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateProfileByID(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var userReq request.UpdateProfile
	err = ctx.ShouldBindJSON(&userReq)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user := model.User{
		ID:       userID,
		Phone:    userReq.Phone,
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	err = h.usersService.UpdateUser(ctx, user)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := response.UpdateProfile{
		ID:        user.ID,
		Phone:     user.Phone,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteProfile(ctx *gin.Context) {
	userID, err := helper.GetCurrentUserIDFromToken(ctx)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	err = h.usersService.DeleteUserByID(ctx, userID)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (h *Handler) DeleteProfileByID(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = h.usersService.DeleteUserByID(ctx, userID)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (h *Handler) GetProfile(ctx *gin.Context) {
	userID, err := helper.GetCurrentUserIDFromToken(ctx)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	user, err := h.usersService.GetUserByID(ctx, userID)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := response.GetProfile{
		ID:        user.ID,
		Role:      user.Role.Name(),
		Phone:     user.Phone,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetProfileByID(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := h.usersService.GetUserByID(ctx, userID)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := response.GetProfile{
		ID:        user.ID,
		Role:      user.Role.Name(),
		Phone:     user.Phone,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, resp)
}
