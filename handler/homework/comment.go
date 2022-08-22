package homework

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/file"
	"github.com/MuXiFresh-be/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type CommentRequest struct {
	HomeworkID uint   `json:"homework_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

// @Summary  Comment on Homework
// @Description Add comment to one's homework
// @Tags homework
// @Accept  json
// @Produce  json
// @Param Authorization header string true "token 用户令牌"
// @Param req body CommentRequest true  "HomeworkID 评论作业的id || Content 评论内容"
// @Success 200 {object} handler.Response
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /homework/comment [post]
func Comment(c *gin.Context) {
	log.Info("Idea getIdeaList function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	email := c.MustGet("email").(string)
	var req CommentRequest
	if err := c.ShouldBind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}
	// 发布评论
	if err := file.CommentHomework(email, req.HomeworkID, req.Content); err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, "success")
}

// @Summary Get comments
// @Description 查看已发布的评论
// @Tags homework
// @Accept  json/application
// @Produce  json/application
// @Param Authorization header string true  "获取email"
// @Param id query integer true "id--帖子的id"
// @Param limit query integer true "limit--偏移量指定开始返回记录之前要跳过的记录数 "
// @Param page  query integer true "page--限制指定要检索的记录数 "
// @Success 200 {object}  handler.Response
// @Router /homework/comment [get]
func GetComments(c *gin.Context) {
	log.Info("Idea getIdeaList function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	id := c.Query("id")
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}
	comments, num, err := file.GetComment(id, limit*page, limit)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, map[string]interface{}{
		"comments": comments,
		"num":      num,
	})
}

// @Summary Delete comment
// @Description 删除用户发布的帖子
// @Tags homework
// @Accept  json/application
// @Produce  json/application
// @Param Authorization header string true  "获取email"
// @Param comment_id path integer true "评论的id"
// @Success 200 {string}  json "{"code":0,"message":"OK","data":{}}"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /homework/comment/:comment_id [delete]
func DeleteComment(c *gin.Context) {
	log.Info("Idea getIdeaList function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	email := c.MustGet("email").(string)
	id := c.Param("comment_id")
	if err := file.Delete(id, email); err != nil {
		SendError(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "Success")
}
