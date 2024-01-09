package router

import (
	"as/handler"
	"as/handler/application"
	"as/handler/car"
	"as/handler/collection"
	"as/handler/sd"
	"as/handler/user"
	"as/pkg/constvar"
	"as/pkg/errno"
	"as/router/middleware"
	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		handler.SendError(c, errno.ErrIncorrectAPIRoute, nil, "", "")
	})
	//
	// swagger API doc
	// g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 权限要求，普通用户/管理员/超管
	normalRequired := middleware.AuthMiddleware(constvar.AuthLevelNormal)
	adminRequired := middleware.AuthMiddleware(constvar.AuthLevelAdmin)

	// auth 模块
	authRouter := g.Group("api/v1/auth")
	{
		authRouter.POST("/login", user.Login)
		authRouter.POST("/register", user.Register)
	}

	// user 模块
	userRouter := g.Group("api/v1/user")
	userRouter.Use(normalRequired)
	{
		userRouter.GET("/profile/:id", user.GetProfile)
		userRouter.GET("/myprofile", user.GetMyProfile)
		userRouter.GET("/list", user.List)
		userRouter.PUT("", user.UpdateInfo)
	}

	carRouter := g.Group("api/v1/car")
	{
		carRouter.POST("", adminRequired, car.Create)
		carRouter.PUT("", adminRequired, car.Update)
		carRouter.DELETE("/:id", adminRequired, car.Delete)

		carRouter.GET("/info/:id", car.Get)
		carRouter.GET("/list", car.List)
	}

	// collection
	collectionRouter := g.Group("api/v1/collection")
	collectionRouter.Use(normalRequired)
	{
		collectionRouter.GET("/list", collection.List)
		collectionRouter.POST("/:car_id", collection.CreateOrRemove)
	}

	applicationRouter := g.Group("api/v1/application")
	applicationRouter.Use(normalRequired)
	{
		applicationRouter.POST("", application.Create)
		applicationRouter.GET("/list", application.ListMyApplication)
		applicationRouter.GET("/list/all", application.ListAllApplication)
		// applicationRouter.GET("/info/:id", application.Get)
		// applicationRouter.PUT("", application.Update)
		// applicationRouter.DELETE("/info/:id", application.Delete)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
