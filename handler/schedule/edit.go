package schedule

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/schedule"
	"github.com/gin-gonic/gin"
)

// @Summary "编辑报名表"
// @Description
// @Tags form
// @Accept application/json
// @Produce application/json
// @Param object body editRequest true "edit_request"
// @Success 200 {object} editResponse
// @Router /form [put]
func Edit(c *gin.Context) {
	var request editRequest
	email := c.MustGet("email").(string)
	if err := c.BindJSON(&request); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	if err := schedule.EditForm(email, request.Name, request.Major, request.Group); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, "编辑报名表成功！")
}
