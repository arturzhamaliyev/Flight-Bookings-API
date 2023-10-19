package handler

import (
	"net/http"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) createUser(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		zap.S().Infof("failed to bind json: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = h.users.CreateUser(ctx, user)
	if err != nil {
		zap.S().Infof("failed to create user: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, "")
}
