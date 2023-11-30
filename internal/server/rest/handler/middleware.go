package handler

import (
	"net/http"
	"strings"

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

func (h *Handler) Auth(ctx *gin.Context) {
	clientToken := ctx.Request.Header.Get("Authorization")
	if clientToken == "" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "no Authorization header provided"})
		return
	}

	extractedToken := strings.Split(clientToken, "Bearer ")
	if len(extractedToken) == 2 {
		clientToken = strings.TrimSpace(extractedToken[1])
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid format of Authorization Token"})
		return
	}

	claims, err := h.sessionService.ValidateToken(clientToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.Set("email", claims.UserEmail)

	ctx.Next()
}
