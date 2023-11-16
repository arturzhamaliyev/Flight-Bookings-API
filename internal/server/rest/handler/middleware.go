package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) IdempotencyCheck(ctx *gin.Context) {
	guid := ctx.GetHeader("X-Idempotency-Key")
	_, err := uuid.Parse(guid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Next()
}
