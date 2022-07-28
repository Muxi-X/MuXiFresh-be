package homework

import (
	"github.com/gin-gonic/gin"
)

type HomeworkRequest struct {
	Content string
}

// PublishHomework ... 发布作业
// @Summary Homework api
// @Description
// @Tags homework
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body
// @Success 200 {object} loginResponse
// @Router /homework [post]
func PublishHomework(c *gin.Context) {

}
