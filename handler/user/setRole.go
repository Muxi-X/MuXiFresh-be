package user

import (
	"fmt"

	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary "设置权限"
// @Description "修改某个用户的权限等级"
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "token 用户令牌"
// @Param object body modifyRoleRequest true "modifyRoleRequest"
// @Success 200
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router api/v1/user/role [put]
func SetRole(c *gin.Context) {
	log.Info("Set one member's role function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	// var ownRole uint32

	ownRole := c.MustGet("role").(uint32)
	fmt.Println(12367, ownRole)
	//	userId := c.MustGet("userId").(uint32)

	var role modifyRoleRequest

	if err := c.Bind(&role); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	if role.Role > ownRole || ownRole < 2 {
		SendResponse(c, nil, "您的权限等级过低，无法进行该操作")
		return
	}

	if err := model.DB.Self.Table("users").Where("email = ? ", role.Email).Updates(map[string]interface{}{"role": role.Role}).Error; err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "权限修改成功")
}
