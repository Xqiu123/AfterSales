package main

import (
	"as/dao"
	"as/log"
	"as/router"
	"as/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @Title as
// @Version 1.0
// @Description The gateway of as
// @Host as.muxixyz.com
// @BasePath /api/v1

// @tag.name user
// @tag.description 用户服务
// @tag.name auth
// @tag.description 认证服务

func main() {
	// logger sync
	defer log.SyncLogger()

	dao.Init()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	// Routes.
	router.Load(
		// Cores.
		g,

		// MiddleWares.
		middleware.Logging(),
		middleware.RequestId(),
	)

	log.Info(g.Run().Error())
}
