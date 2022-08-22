package form

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/form"
	"github.com/gin-gonic/gin"
)

// @Summary "查看报名表（分组）"
// @Description
// @Tags form
// @Accept application/json
// @Produce application/json
// @Param object body searchRequest true "按组查询报名表"
// @Success 200
// @Router /form/search [post]
func Search(c *gin.Context) {
	email := c.MustGet("email").(string)
	if email == "" {
		return
	}
	var group searchRequest
	if err := c.BindJSON(&group); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	Form, err := form.SearchForm(group.Group)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, Form)
}
