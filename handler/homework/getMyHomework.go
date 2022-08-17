package homework

import "github.com/gin-gonic/gin"

// @Summary Get My Homework
// @Tags homework
// @Description 查看我的作业
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param group_id query integer true "group_id--小组id 1-设计组 2-产品组 3-安卓组 4-前端组 5-后端组"
// @Success 200 {object} []file.HomeworkPublished {} "{"msg":"查看成功"}"
// @Failure 500 {object} errno.Errno "{"msg":"Error occurred while getting url queries."}"
// @Router /homework/published/:id/mine [get]
func GetMyHomework(c *gin.Context) {
	email := c.MustGet("email").(string)
	gropID := c.Param("group_id")
	homework :=
}
