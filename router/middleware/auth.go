package middleware

import (
	"github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/auth"
	"github.com/MuXiFresh-be/pkg/errno"

	"github.com/gin-gonic/gin"
)

// 用户权限鉴定
type UserAuth struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Role      uint32 `json:"role"`
	StudentID string `json:"student_id"`
}

// AuthMiddleware ... 认证中间件
// limit 为限制的权限等级 //以下为对权限等级的限制
func AuthMiddleware(limit uint32) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		ctx, err := auth.ParseRequest(c)
		if err != nil {
			handler.SendResponse(c, errno.ErrAuthToken, err.Error())
			c.Abort()
			return
		}

		c.Set("userId", ctx.Id)
		c.Set("email", ctx.Email)
		c.Set("role", ctx.Role)
		c.Set("expiresAt", ctx.ExpiresAt)

		c.Next()
	}
}
