package user

import (
	"fmt"
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/pkg/errno"
	service "github.com/MuXiFresh-be/service/user"
	"github.com/MuXiFresh-be/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary Update User Info
// @Description 更改用户信息
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true  "token 用户令牌"
// @Param req body updateInfoRequest true  "Avatar头像|| NickName昵称"
// @Success 200  "Success"
// @Failure 400 {string} json  "{"Code":400, "Message":"Error occurred while binding the request body to the struct","Data":nil}"
// @Failure 500 {string} json  "{"Code":500,"Message":"Database error","Data":nil}"
// @Router /user [put]
func UpdateInfo(c *gin.Context) {
	log.Info("student login function called.", zap.String("X-Request-Id", util.GetReqID(c)))
	var req updateInfoRequest
	if err := c.ShouldBind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	fmt.Println("------req:", req, "---------")
	if req.Email == "" {
		req.Email = c.MustGet("email").(string)
	}
	if err := service.UpdateInfo(req.Email, req.AvatarURL, req.Name); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, "Success")
}

