package auth

import (
	"fmt"
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/handler/user"
	"github.com/MuXiFresh-be/pkg/errno"
	service "github.com/MuXiFresh-be/service/user"
	"github.com/gin-gonic/gin"
)

// @Summary Register
// @Tags auth
// @Description 邮箱注册登录
// @Accept application/json
// @Produce application/json
// @Param object body auth.CreateUserRequest true "注册用户信息"
// @Success 200 {object} handler.Response "{"msg":"将student_id作为token保留"}"
// @Failure 401 {object} errno.Errno "{"error_code":"10001", "message":"The email address has been registered"} "
// @Failure 400 {object} errno.Errno "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} errno.Errno "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var req user.RegisterRequest

	if err := c.ShouldBind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine()) //, errno.ErrBind)
		fmt.Println(req)
		return
	}

	err := service.Register(req)

	if err != nil {
		SendBadRequest(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "succeed in registration")

}
