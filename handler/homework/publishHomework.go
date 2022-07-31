package homework

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	Service "github.com/MuXiFresh-be/service/file"
	"github.com/gin-gonic/gin"
)

// PublishHomework ... 发布作业
// @Summary Homework api
// @Description
// @Tags homework
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body HomeworkRequest  true "content--作业内容 ; group_id--小组id   1-设计组 2-产品组 3-安卓组 4-前端组 5-后端组"
// @Success 200 {object} loginResponse
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /homework [post]
func PublishHomework(c *gin.Context) {
	var homework HomeworkRequest
	email := c.MustGet("email").(string)
	if err := c.ShouldBind(&homework); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}
	if err := Service.PublishHomework(email, homework.GroupID, homework.Content); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, "Success")

}
