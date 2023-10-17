package http

import (
	"net/http"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/errors"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func handleError(ctx *gin.Context, err error) {
	logging.From(ctx).Info("error occured in request", zap.Error(err))

	var code int

	switch {
	case errors.Is(err, errors.ErrValidation):
		code = http.StatusBadRequest
	case errors.Is(err, errors.ErrNotFound):
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}

	ctx.JSON(code, err)
}
