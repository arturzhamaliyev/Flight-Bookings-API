package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest/handler/request"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest/handler/response"
)

func (h *Handler) AddAdmin(ctx *gin.Context) {
	var userReq request.SignUp
	err := ctx.ShouldBindJSON(&userReq)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user := model.User{
		ID:        uuid.New(),
		Role:      model.Admin,
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
