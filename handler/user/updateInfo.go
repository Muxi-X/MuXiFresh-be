package user

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/pkg/errno"
	service "github.com/MuXiFresh-be/service/user"
	"github.com/MuXiFresh-be/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary 更改用户信息
// @Description 更改用户信息
// @Tags user
// @Accept  json
// @Produce  json
// @Param  token header string true  "UserId--用户ID"
// @Param req body updateInfoRequest true  "Avatar头像|| NickName昵称"
// @Success  200 "成功"
// @Router /api/v1/user [put]
// UpdateUserInfo ... 更改用户的信息
func UpdateInfo(c *gin.Context) {
	log.Info("student login function called.", zap.String("X-Request-Id", util.GetReqID(c)))
	var req updateInfoRequest
	if err := c.ShouldBind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}
	if err := service.UpdateInfo(req.Email, req.Email, req.AvatarURL); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
}
