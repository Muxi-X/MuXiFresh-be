package form

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	F "github.com/MuXiFresh-be/service/form"
	U "github.com/MuXiFresh-be/service/user"
	"github.com/gin-gonic/gin"
)

// @Summary "创建报名表"
// @Description
// @Tags form
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body createRequest true "create_request"
// @Success 200
// @Router /form [post]
func Create(c *gin.Context) {
	var request createRequest
	email := c.MustGet("email").(string)
	if err := c.BindJSON(&request); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	userForm, err := F.CreateForm(email,
		request.Group, request.Reason, request.Understand,
		request.SelfIntroduction, request.IfOtherOrganization,request.Work)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}
	if err := U.UpdateInfor(email, request.Avatar, request.Name, request.StudentId, request.College, request.Major, request.Grade, request.Gender,  request.PhoneNumber, request.QqNumber); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	if err := F.UpdateSchedule(email, request.Group); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, userForm)
}
