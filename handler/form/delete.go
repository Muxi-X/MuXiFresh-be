package form

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/form"
	"github.com/gin-gonic/gin"
)

// @Summary "删除报名表、进度等信息"
// @Description 
// @Tags form
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param object body deleteRequest true "delete_request"
// @Success 200
// @Router /form/delete [post]
func Delete(c *gin.Context) {
	email := c.MustGet("email").(string)
	if email == "" {
		return
	}

	var request deleteRequest
	if err := c.BindJSON(&request); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	err := form.DeleteForm(request.Email)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	err = form.DeleteSchedule(request.Email)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "Delete success")
}