package form

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/gin-gonic/gin"
)

// @Summary "编辑报名表"
// @Description
// @Tags form
// @Accept application/json
// @Produce application/json
// @Param object body editRequest true "edit_request"
// @Success 200 {object} editResponse
// @Router /form/edit [post]

func Edit(c *gin.Context) {
	var form editRequest
	email := c.MustGet("email").(string)
	if err := c.BindJSON(&form); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	if err := model.DB.Self.Table("forms").Where("email=?", email).Updates(map[string]interface{}{"name": form.Name, "avatar": form.Avatar, "student_id": form.StudentId, "college": form.College, "major": form.Major, "grade": form.Grade, "gender": form.Gender, "contact_way": form.Contact_way, "contact_number": form.Contact_number, "group": form.Group, "reason": form.Reason, "understand":form.Understand, "self_introduction":form.Self_introduction, "if_other_organization":form.If_other_organization}).Error; err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, "编辑报名表成功！")
}
