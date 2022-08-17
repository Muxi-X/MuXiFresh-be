package form

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	Form "github.com/MuXiFresh-be/service/form"
	"github.com/gin-gonic/gin"
)

// @Summary "创建报名表"
// @Description
// @Tags form
// @Accept application/json
// @Produce application/json
// @Param object body createRequest true "create_request"
// @Success 200 {object} form.FormModel "{"msg":"创建报名表成功"}"
// @Router /form [post]
func Create(c *gin.Context) {
	var request createRequest
	email := c.MustGet("email").(string)
	if err := c.BindJSON(&request); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	userForm, err := Form.CreateForm(email, request.Name, request.Avatar, request.StudentId,
		request.College, request.Major, request.Grade,
		request.Gender, request.ContactWay, request.ContactNumber,
		request.Group, request.Reason, request.Understand,
		request.SelfIntroduction, request.IfOtherOrganization)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, userForm)
}
