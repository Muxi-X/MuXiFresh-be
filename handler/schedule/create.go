package schedule

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/schedule"
	"github.com/gin-gonic/gin"
)

// @Summary "创建报名表"
// @Description
// @Tags form
// @Accept application/json
// @Produce application/json
// @Param object body createRequest true "create_request"
// @Success 200 {object} handler.Response "{"msg":"创建报名表成功"}"
// @Router /form [post]
func Create(c *gin.Context) {
	var request createRequest
	email := c.MustGet("email").(string)
	if err := c.BindJSON(&request); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	if err := schedule.CreateForm(email, request.Name, request.StudentId, request.Major, request.Group); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, "创建报名表成功！")
}
