package form

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/form"
	"github.com/gin-gonic/gin"
)

// @Summary "查看报名表（其他成员）"
// @Description 通过email查看成员报名表
// @Tags form
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param object body viewothersRequest true "查看他人报名表"
// @Success 200
// @Router /form/view [post]
func ViewOthers(c *gin.Context) {
	email := c.MustGet("email").(string)
	if email == "" {
		return
	}

	var request viewothersRequest
	if err := c.BindJSON(&request); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	Form, err := form.ViewForm(request.Email)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, Form)
}