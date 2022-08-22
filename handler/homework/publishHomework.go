package homework

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/pkg/errno"
	Service "github.com/MuXiFresh-be/service/file"
	"github.com/MuXiFresh-be/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// PublishHomework ... 发布作业
// @Summary Publish homework
// @Description
// @Tags homework
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body HomeworkRequest  true "content--作业内容 ; group_id--小组id   1-设计组 2-产品组 3-安卓组 4-前端组 5-后端组"
// @Success 200 {object} handler.Response
// @Failure 400 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /homework/publish [post]
func PublishHomework(c *gin.Context) {
	log.Info("Idea getIdeaList function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	var req PublishHomeworkRequest

	email := c.MustGet("email").(string)

	if err := c.ShouldBind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	homework, err := Service.PublishHomework(email, req.GroupID, req.Title, req.Content, req.FileUrl)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, homework)

}
