package form

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/model/form"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/gin-gonic/gin"
)

// @Summary "查看报名表"
// @Description
// @Tags form
// @Accept application/json
// @Produce application/json
// @Success 200 {object} createResponse
// @Router /form/view [get]
func View(c *gin.Context) {
	email := c.MustGet("email").(string)
	Form, err := form.ViewForm(email)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, Form)
}
