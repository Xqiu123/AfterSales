package main

import (
	"as/config"
	"as/dao"
	"as/log"
	"as/router"
	"as/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @Title as
// @Version 1.0
// @Description The gateway of AfterSales
// @Host localhost
// @BasePath /api/v1

func main() {
	err := config.Init("./conf/config.yaml", "")
	if err != nil {
		panic(err)
	}

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
