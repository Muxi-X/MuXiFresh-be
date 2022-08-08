package user

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/pkg/errno"

	service "github.com/MuXiFresh-be/service/user"

	"github.com/MuXiFresh-be/util"
	"strconv"
	// pb "forum-user/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetProfile ... 获取 userProfile
// @Summary Get user_profile api
// @Description 通过 userId 获取完整 user 信息
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param id path int true "user_id"
// @Success 200 {object} userProfile
// @Failure 400 {string} json  "{"Code":400,"Message":"Error occurred while getting path param.","Data":nil}"
// @Failure 500 {string} json  "{"Code":500,"Message":"Internal server error","Data":nil}"
// @Router /user/profile/:id [get]
func GetProfile(c *gin.Context) {
	log.Info("User getInfo function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	var id int
	var err error

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 发送请求
	getProfileResp, err := service.GetProfile(id)

	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回 response
	user := &userProfile{
		Id:     getProfileResp.ID,
		Name:   getProfileResp.Name,
		Avatar: getProfileResp.Avatar,
		Email:  getProfileResp.Email,
		Role:   getProfileResp.Role,
	}

	//if err != nil {
	//	// TODO: 判断错误是否是用户不存在
	//	SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
	//	return
	//}

	SendResponse(c, nil, user)
}
