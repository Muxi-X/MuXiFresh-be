package user

import (
	"fmt"
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	// pb "forum-user/proto"
	"strconv"
)

// List ... 获取 userList
// @Summary get user_list api
// @Description 通过 group 和 team 获取 user_list
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Param last_id query int false "last_id"
// @Param Authorization header string true "token 用户令牌"
// @Param group_id path int true "group_id"
// @Param team_id path int true "team_id"
// @Success 200 {object} listResponse
// @Router /user/list [get]
func List(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	log.Info("User getInfo function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 limit
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 获取 page
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	lastId, err := strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}
	fmt.Println(limit, page, lastId)

	// 构造请求给 list
	// listReq := &pb.ListRequest{
	// 	LastId: uint32(lastId),
	// 	Offset: uint32(page * limit),
	// 	Limit:  uint32(limit),
	// }

	// 发送请求
	// listResp, err := service.UserClient.List(context.Background(), listReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// SendResponse(c, nil, listResp)
}
