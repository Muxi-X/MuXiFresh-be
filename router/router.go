package router

import (
	"github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/handler/auth"
	"github.com/MuXiFresh-be/handler/sd"
	"github.com/MuXiFresh-be/pkg/constvar"
	"github.com/MuXiFresh-be/pkg/errno"

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
		handler.SendError(c, errno.ErrIncorrectAPIRoute, nil, "", "")
	})

	// swagger API doc
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	normalRequired := middleware.AuthMiddleware(constvar.AuthLevelNormal)
	adminRequired := middleware.AuthMiddleware(constvar.AuthLevelAdmin)
	superAdminRequired := middleware.AuthMiddleware(constvar.AuthLevelSuperAdmin)

	// auth
	authRouter := g.Group("api/v1/auth")
	{
		authRouter.POST("/register", auth.Register)

		authRouter.PUT("/authorize/:email/:role", superAdminRequired, auth.Authorize)

		authRouter.GET("/administrator", normalRequired, auth.GetAdministrator)
	}

	// user 模块
	userRouter := g.Group("api/v1/user")
	{
		userRouter.POST("/login", user.Login)

		userRouter.PUT("", normalRequired, user.UpdateInfo)

		userRouter.GET("/info", normalRequired, user.GetInfo)

		userRouter.GET("/profile/:email", user.GetProfile)

		//userRouter.GET("/list", user.List)

		userRouter.GET("/qiniu_token", user.GetQiniuToken)

	}

	// homework 模块
	homework := g.Group("api/v1/homework").Use(middleware.AuthMiddleware(constvar.AuthLevelNormal))
	{
		// 管理员发布作业
		homework.POST("/publish", adminRequired, Homework.PublishHomework)

		// 获取已经发布的作业
		homework.GET("/published", Homework.GetHomeworkPublished)

		// 获取已发布作业的详细内容
		homework.GET("/published/details/:id", Homework.GetHomeworkDetails)

		// 获取作业详细
		homework.GET("/review", adminRequired, Homework.ReviewHomework)

		// 按小组，获取已经提交的作业
		homework.GET("/handed", adminRequired, Homework.GetHandedHomework)

		// 评论作业
		homework.POST("/comment", adminRequired, Homework.Comment)

		// 获取某个作业的全部评论
		homework.GET("/comment", adminRequired, Homework.GetComments)

		// 删除评论
		homework.DELETE("/comment/:comment_id", adminRequired, Homework.DeleteComment)

		// 上传作业
		homework.POST("", normalRequired, Homework.UploadHomework)

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
