package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"oasis/ready/biz/handlers"
	"oasis/ready/docs"
	"oasis/ready/middle_ware"
)

func register(r *gin.Engine) {
	r.Use(middle_ware.CustomRecovery(), //错误恢复堆栈现场
		middle_ware.Cors(),           //跨域设置
		middle_ware.CheckRequestId()) //填充请求Id


	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group("/api/v1")
	v1.GET("/ping", handlers.Ping)
}
