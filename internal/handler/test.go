package handler

import "github.com/gin-gonic/gin"

func (h *Handler) Test(ctx *gin.Context) {
	resp := h.service.Test()
	ctx.JSON(resp.Code, resp)
}
