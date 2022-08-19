package form

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/form"
	U "github.com/MuXiFresh-be/service/user"
	"github.com/gin-gonic/gin"
)

// @Summary "编辑报名表"
// @Description
// @Tags form
// @Accept application/json
// @Produce application/json
// @Param object body editRequest true "edit_request"
// @Success 200 {object} form.FormModel
// @Router /form [put]
func Edit(c *gin.Context) {
	var request editRequest
	email := c.MustGet("email").(string)
	if err := c.BindJSON(&request); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	err := form.EditForm(email, request.College, request.Major,
		request.Grade, request.Gender, request.ContactWay,
		request.ContactNumber, request.Group, request.Reason,
		request.Understand, request.SelfIntroduction, request.IfOtherOrganization)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	if err := U.UpdateInfor(email, request.Avatar, request.Name,request.StudentId); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "Success")
}
