package homework

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/service/file"
	"github.com/gin-gonic/gin"
)

// @Summary Get My Homework
// @Tags homework
// @Description 查看我的作业
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param id path integer true "homework 的id"
// @Success 200 {object} []file.Homework {} "{"msg":"查看成功"}"
// @Failure 500 {object} errno.Errno "{"msg":"Error occurred while getting url queries."}"
// @Router /homework/published/:id/mine [get]
func GetMyHomework(c *gin.Context) {
	email := c.MustGet("email").(string)
	id := c.Param("id")
	homework, err := file.GetMine(email, id)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, homework)
}
