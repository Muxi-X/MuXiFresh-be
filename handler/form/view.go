package form

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/form"
	U "github.com/MuXiFresh-be/service/user"
	"github.com/gin-gonic/gin"
)

// @Summary "查看报名表（个人）"
// @Description
// @Tags form
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param Authorization header string true "token"
// @Success 200
// @Router /form/view [get]
func View(c *gin.Context) {
	email := c.MustGet("email").(string)
	Form, err := form.ViewForm(email)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	User, err := U.GetInfo(email)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	var response = viewResponse{
		StudentId:           User.StudentId,
		Email:               User.Email,
		Name:                User.Name,
		Avatar:              User.Avatar,
		College:             User.College,
		Major:               User.Major,
		Grade:               User.Grade,
		Gender:              User.Gender,
		ContactWay:          User.ContactWay,
		ContactNumber:       User.ContactNumber,
		Group:               Form.Group,
		Reason:              Form.Reason,
		Understand:          Form.Understand,
		SelfIntroduction:    Form.SelfIntroduction,
		IfOtherOrganization: Form.IfOtherOrganization,
	}

	//SendResponse(c, nil, Form)
	SendResponse(c, nil, response)
}
