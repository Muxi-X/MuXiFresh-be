package router

import (
	"github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/handler/auth"
	"github.com/MuXiFresh-be/handler/schedule"
	"github.com/MuXiFresh-be/handler/sd"
	"github.com/MuXiFresh-be/handler/user"
	"github.com/MuXiFresh-be/pkg/constvar"
	"github.com/MuXiFresh-be/pkg/errno"

	Form "github.com/MuXiFresh-be/handler/form"
	Homework "github.com/MuXiFresh-be/handler/homework"
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

		// userRouter.GET("/list", user.List)

		userRouter.GET("/qiniu_token", user.GetQiniuToken)

		userRouter.PUT("/role", normalRequired, user.SetRole)

	}

	// schedule 模块
	scheduleRouter := g.Group("api/v1/schedule").Use(normalRequired) // 设置中间件，并确定用户等级
	{

		scheduleRouter.GET("", schedule.ViewOwnSchedule)

		scheduleRouter.PUT("/admit", adminRequired, schedule.Admit)

		scheduleRouter.PUT("/cancel_admission/:email", adminRequired, schedule.CancelAdmission)
	}

	// homework 模块
	homework := g.Group("api/v1/homework").Use(middleware.AuthMiddleware(constvar.AuthLevelNormal))
	{
		// 管理员发布作业
		homework.POST("/publish", adminRequired, Homework.PublishHomework)

		// 获取已经发布的作业
		homework.GET("/published", Homework.GetHomeworkPublished)

		// 获取我写的作业
		homework.GET("/published/:id/mine", Homework.GetMyHomework)

		// 查看我所有作业的完成状况
		homework.GET("/performance", Homework.GetMyPerformance)

		// 获取已发布作业的详细内容
		homework.GET("/published/details/:id", Homework.GetHomeworkDetails)

		// 获取作业详细
		homework.GET("/review", adminRequired, Homework.ReviewHomework)

		// 获取某人的全部作业
		homework.GET("/:email", adminRequired, Homework.GetOtherHomework)

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

		// 修改作业
		homework.POST("/change/uploaded/:id", normalRequired, Homework.ModifyHomework)

		// 修改发布的作业
		homework.POST("/change/published/:id", adminRequired, Homework.ModifyPublished)

	}

	// form 模块
	formRouter := g.Group("api/v1/form").Use(middleware.AuthMiddleware(constvar.AuthLevelNormal))
	{
		// 创建报名表
		formRouter.POST("", Form.Create)
		// 修改报名表
		formRouter.PUT("", Form.Edit)
		// 查看报名表（个人）
		formRouter.GET("/view", Form.View)
		// 分组查询成员信息
		formRouter.POST("/search", Form.Search)
		// 查看报名表（他人）
		formRouter.POST("/view", Form.ViewOthers)
		// 删除报名表和进度信息
		formRouter.POST("/delete", Form.Delete)
		// 移动分组
		formRouter.POST("/movegroup", Form.MoveGroup)
	}

	// The health check Fandlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
