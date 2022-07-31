package homework

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/file"
	"github.com/MuXiFresh-be/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CommentRequest struct {
	HomeworkID uint   `json:"homework_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

// @Summary  comment
// @Description Add comment to one's homework
// @Tags homework
// @Accept  json
// @Produce  json
// @Param Authorization header string true "token 用户令牌"
// @Param req body CommentRequest true  "HomeworkID 评论作业的id || Content 评论内容"
// @Success 200 "成功"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /homework/comment [post]
func Comment(c *gin.Context) {
	email := c.MustGet("email").(string)
	var req CommentRequest
	if err := c.ShouldBind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}
	// 发布评论
	if err := file.CommentHomework(email, req.HomeworkID, req.Content); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, "Success")
}

// @Summary delete comment
// @Description 删除用户发布的帖子
// @Tags forum
// @Accept  json/application
// @Produce  json/application
// @Param Authorization header string true  "获取email"
// @Param id path string true "id--评论的ID"
// @Success 200 {string}  json "{"code":0,"message":"OK","data":{}}"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /forum/comment/:comment_id [delete]
func DeleteComment(c *gin.Context) {
	log.Info("Idea getIdeaList function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	email := c.MustGet("email").(string)
	id := c.PostForm("id")
	if err := file.Delete(id, email); err != nil {
		SendError(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "Success")
}
