package router

import (
	"github.com/MuXiFresh-be/handler/auth"
	"github.com/MuXiFresh-be/handler/sd"
	"github.com/MuXiFresh-be/pkg/constvar"
	"net/http"

	_ "github.com/MuXiFresh-be/docs"
	Homework "github.com/MuXiFresh-be/handler/homework"
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

	normalRequired := middleware.AuthMiddleware(constvar.AuthLevelNormal)
	adminRequired := middleware.AuthMiddleware(constvar.AuthLevelAdmin)
	// superAdminRequired := middleware.AuthMiddleware(constvar.AuthLevelSuperAdmin)

	// auth
	authRouter := g.Group("api/v1/auth")
	{
		authRouter.POST("/register", auth.Register)
	}

	// user 模块
	userRouter := g.Group("api/v1/user")
	{
		userRouter.POST("/login", user.Login)
		userRouter.GET("/profile/:id", normalRequired, user.GetProfile)
		userRouter.GET("/list", user.List)
		userRouter.POST("/avatar", user.UploadAvatar)

	}

	// homework 模块
	homework := g.Group("api/v1/homework").Use(middleware.AuthMiddleware(constvar.AuthLevelNormal))
	{
		homework.POST("/publish", adminRequired, Homework.PublishHomework)

		homework.POST("", Homework.UploadHomework)
		//homework.GET("", Homework.DownloadHomework)

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
