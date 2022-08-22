package homework

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/file"
	"github.com/gin-gonic/gin"
)

// ModifyPublished ... 修改作业
// @Summary Modify published homework
// @Description 修改已发布的作业
// @Tags homework
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body homework.ModifyPublishedHomeworkRequest true "修改内容"
// @Param id path integer true "发布的的作业id"
// @Success 200 {object} handler.Response
// @Failure 400 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /homework/change/published/:id [post]
func ModifyPublished(c *gin.Context) {
	email := c.MustGet("email").(string)
	ID := c.Param("id")
	var req ModifyPublishedHomeworkRequest
	if err := c.ShouldBind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	if err := file.ModifyPublished(ID, email, req.GroupID, req.Title, req.Content, req.FileUrl); err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, "Success")
}
