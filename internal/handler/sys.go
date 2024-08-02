package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health		 godoc
// @Summary      Return health status
// @Tags		 System
// @Produce      json
// @Success		 200 {object} models.StatusResponse
// @Router       /market/health [get]
func (h *Handler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
