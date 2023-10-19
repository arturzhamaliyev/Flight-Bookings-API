package http

import (
	"net/http"

	e "github.com/arturzhamaliyev/Flight-Bookings-API/internal/errors"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/users/model"
	"github.com/gin-gonic/gin"
)

func (s *Server) createUser(ctx *gin.Context) {
	// ctx.Header("Content-Type", "application/json")

	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		s.logger.Infof("failed to bind json: %v", err)
		handleError(ctx, e.Wrap(e.ErrInvalidRequest, err))
		return
	}

	err = s.users.CreateUser(ctx, user)
	if err != nil {
		s.logger.Infof("failed to create user: %v", err)
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, "")
}
