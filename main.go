package main

import (
	"github.com/gin-gonic/gin"

	"oasis/ready/global"
	"oasis/ready/initializer"
)

func init() {
	// 初始化配置
	initializer.InitializeConfig()

	// 初始化日志设置
	global.App.Log = initializer.InitializeLog()
	global.App.Log.Info("log init success!")

	// 初始化数据库
	global.App.DB = initializer.InitializeDB()
	global.App.Log.Info("db init success!")

	// 初始化redis
	global.App.Redis = initializer.InitializeRedis()
	global.App.Log.Info("redis init success!")

}

// @title Ready
// @version 1.0
// @description Ready API
// @BasePath /api/v1
// @query.collection.format multi
func main() {

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	r := gin.Default()

	register(r)

	r.Run(":" + global.App.Config.App.Port)

}

