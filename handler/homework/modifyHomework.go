package homework

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/file"
	"github.com/gin-gonic/gin"
)

// ModifyHomework ... 修改作业
// @Summary Modify homework
// @Description 修改已经上传的作业
// @Tags homework
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body homework.ModifyHomeworkRequest true "修改内容"
// @Param id path integer true "上传的作业id"
// @Success 200 {object} handler.Response
// @Failure 400 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /homework/change/uploaded/:id [post]
func ModifyHomework(c *gin.Context) {
	ID := c.Param("id")
	email := c.MustGet("email").(string)
	var req ModifyHomeworkRequest
	if err := c.ShouldBind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}
	if err := file.ModifyHW(ID, email, req.Title, req.Content, req.FileUrl); err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "Success")
}
