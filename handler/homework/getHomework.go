package homework

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/file"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary Get Homework
// @Tags homework
// @Description 查看不同组别的作业
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param group_id query integer true "group_id--小组id 1-设计组 2-产品组 3-安卓组 4-前端组 5-后端组"
// @Param limit query integer true "limit--偏移量指定开始返回记录之前要跳过的记录数 "
// @Param page  query integer true "page--限制指定要检索的记录数 "
// @Success 200 {object} []file.HomeworkPublished {} "{"msg":"查看成功"}"
// @Failure 500 {object} errno.Errno "{"msg":"Error occurred while getting url queries."}"
// @Router /homework/published [get]
func GetHomeworkPublished(c *gin.Context) {
	ID, err := strconv.Atoi(c.DefaultQuery("group_id", "1"))
	if err != nil {
		SendError(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}
	var limit, page int

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	homework, num, err := file.GetHomework(ID, limit*page, limit)

	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, map[string]interface{}{
		"homework": homework,
		"num":      num,
	})
}
