package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	// ErrNotFound is for errors that occured if result is not found.
	ErrNotFound = errors.New(http.StatusText(http.StatusNotFound))
	// ErrInternalServer is for errors that occured if something happened in application.
	ErrInternalServer = errors.New(http.StatusText(http.StatusInternalServerError))
)

func (h *Handler) ErrorHandler(ctx *gin.Context) {
	ctx.Next()

	for _, err := range ctx.Errors {
		switch err.Err {
		case ErrNotFound:
			ctx.JSON(-1, gin.H{"error": ErrNotFound.Error()})
		default:
			ctx.JSON(-1, gin.H{"error": ErrInternalServer.Error()})
		}
	}
}
