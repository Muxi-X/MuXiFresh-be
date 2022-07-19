package user

import (
	"fmt"
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/util"
	"strconv"
	// pb "forum-user/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetProfile ... 获取 userProfile
// @Summary get user_profile api
// @Description 通过 userId 获取完整 user 信息
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param id path int true "user_id"
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {object} userProfile
// @Router /user/profile/{id} [get]
func GetProfile(c *gin.Context) {
	log.Info("User getInfo function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	var id int
	var err error

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}
	fmt.Println(id)

	// getProfileReq := &pb.GetRequest{Id: id}
	//
	// // 发送请求
	// getProfileResp, err := service.UserClient.GetProfile(context.Background(), getProfileReq)
	// if err != nil {
	// 	SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
	// 	return
	// }
	//
	// // 构造返回 response
	// user := &userProfile{
	// 	Id:     getProfileResp.Id,
	// 	Name:   getProfileResp.Name,
	// 	Avatar: getProfileResp.Avatar,
	// 	Email:  getProfileResp.Email,
	// 	Role:   getProfileResp.Role,
	// }
	//
	// if err != nil {
	// 	// TODO: 判断错误是否是用户不存在
	// 	SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
	// 	return
	// }
	//
	// SendResponse(c, nil, user)
}
