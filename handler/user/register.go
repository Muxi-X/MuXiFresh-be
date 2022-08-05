package user

import (
	MD5 "crypto/md5"
	"encoding/hex"
	"fmt"

	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/handler/schedule"
	User "github.com/MuXiFresh-be/model/user"
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
	var req RegisterRequest

	if err := c.ShouldBind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine()) //, errno.ErrBind)
		fmt.Println(req)
		return
	}

	user := User.UserModel{
		Name:      req.Name,
		StudentId: req.StudentId,
		Email:     req.Email,
		Role:      1,
	}

	md5 := MD5.New()
	md5.Write([]byte(req.Password))
	user.HashPassword = hex.EncodeToString(md5.Sum(nil))

	if err := service.Register(user); err != nil {
		SendBadRequest(c, errno.ErrDatabase, nil, err.Error(), GetLine())
	}
	// if err := user.CreateUser(); err != nil {
	// 	SendBadRequest(c, errno.ErrDatabase, nil, err.Error(), GetLine())
	// }

	//注册成功自动生成进度表
	var sche schedule.Schedules
	sche.AdmissionStatus = 0
	sche.Email = req.Email
	sche.FormStatus = 0
	sche.Name = req.Name
	sche.StudentId = req.StudentId
	sche.WorkStatus = 0
	if err := service.Create(sche); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, "succeed in registration")
}
