package http

import (
	"errors"
	"net/http"

	e "github.com/arturzhamaliyev/Flight-Bookings-API/internal/errors"
	"github.com/gin-gonic/gin"
)

func handleError(ctx *gin.Context, err error) {
	var code int

	switch {
	case errors.Is(err, e.ErrValidation):
		code = http.StatusBadRequest
	case errors.Is(err, e.ErrNotFound):
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}

	ctx.JSON(code, err)
}
