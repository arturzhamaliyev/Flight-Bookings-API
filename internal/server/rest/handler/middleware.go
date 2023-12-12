package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/helper"
)

func (h *Handler) JWTAuthAdmin(ctx *gin.Context) {
	token, err := helper.GetToken(ctx)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}

	err = helper.ValidateToken(token)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		ctx.Abort()
		return
	}

	err = helper.ValidateAdminRole(token)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		ctx.Abort()
		return
	}

	ctx.Next()
}

func (h *Handler) JWTAuthCustomer(ctx *gin.Context) {
	token, err := helper.GetToken(ctx)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}

	err = helper.ValidateToken(token)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		ctx.Abort()
		return
	}

	err = helper.ValidateCustomerRole(token)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		ctx.Abort()
		return
	}

	ctx.Next()
}
