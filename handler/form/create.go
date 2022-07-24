package form

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/model/form"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/gin-gonic/gin"
)

// @Summary "创建报名表"
// @Description
// @Tags form
// @Accept application/json
// @Produce application/json
// @Param object body createRequest true "create_request"
// @Success 200 {object} createResponse
// @Router /form/create [post]
func Create(c *gin.Context) {
	var request createRequest
	email := c.MustGet("email").(string)
	if err := c.BindJSON(&request); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}
	form := form.FormModel{
		Email:                 email,
		Name:                  request.Name,
		StudentId:             request.StudentId,
		College:               request.College,
		Major:                 request.Major,
		Grade:                 request.Grade,
		Gender:                request.Gender,
		Contact_way:           request.Contact_way,
		Contact_number:        request.Contact_number,
		Group:                 request.Group,
		Reason:                request.Reason,
		Understand:            request.Understand,
		Self_introduction:     request.Self_introduction,
		If_other_organization: request.If_other_organization,
	}
	if err := model.DB.Self.Create(&form).Error; err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, "创建报名表成功！")
}
