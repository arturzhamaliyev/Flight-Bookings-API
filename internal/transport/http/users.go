package http

import (
	"net/http"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/errors"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/logging"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/users/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Server) createUser(ctx *gin.Context) {
	// ctx.Header("Content-Type", "application/json")

	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		logging.From(ctx).Info("failed to bind json", zap.Error(err))
		handleError(ctx, errors.ErrInvalidRequest.Wrap(err))
		return
	}

	err = s.users.CreateUser(ctx, user)
	if err != nil {
		logging.From(ctx).Info("failed to create user", zap.Error(err))
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, "")
}
