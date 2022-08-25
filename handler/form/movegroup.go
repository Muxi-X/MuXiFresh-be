package form

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	F "github.com/MuXiFresh-be/service/form"
	"github.com/gin-gonic/gin"
)

// @Summary "移动分组"
// @Description 将email对应的成语移动到指定分组
// @Tags form
// @Accept application/json
// @Produce application/json
// @Param object body moveRequest true "move_request"
// @Success 200 
// @Router /movegroup [post]
func MoveGroup(c *gin.Context) {
	email := c.MustGet("email").(string)
	if email == "" {
		return
	}

	var request moveRequset
	if err := c.BindJSON(&request); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	if err := F.MoveGroup(request.Email,request.Group);err!=nil{
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "Move group success")
}