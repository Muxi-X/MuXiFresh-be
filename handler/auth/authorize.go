package auth

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/user"
	"github.com/MuXiFresh-be/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// @Summary Authorize
// @Description  Modifying User Rights
// @Tags auth
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true  "token 用户令牌"
// @Param id header int true  "被修改用户的id"
// @Param role header int true "权限等级"
// @Success  200 "成功"
// @Failure 500 "失败"
// @Router /auth/authorize/:id/:role [put]
func Authorize(c *gin.Context) {
	log.Info("student login function called.", zap.String("X-Request-Id", util.GetReqID(c)))
	id, _ := strconv.Atoi(c.Param("id"))
	role, _ := strconv.Atoi(c.Param("role"))
	if err := service.Authorize(id, role); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, "Success")
}
