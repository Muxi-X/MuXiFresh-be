package router

import (
	"github.com/MuXiFresh-be/handler/sd"
	"net/http"

	_ "github.com/MuXiFresh-be/docs"
	"github.com/MuXiFresh-be/handler/user"
	"github.com/MuXiFresh-be/router/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
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
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// swagger API doc
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// user 模块
	userRouter := g.Group("api/v1/user")
	{
		userRouter.POST("/login", user.Login)
		userRouter.GET("/profile/:id", user.GetProfile)
		userRouter.GET("/list", user.List)
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
