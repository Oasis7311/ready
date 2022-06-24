package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping
// @Summary Ping接口
// @Schemes
// @Description 测试服务是否连通
// @Tags test
// @Accept json
// @Produce json
// @Router /ping [get]
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "OK",
	})
}
