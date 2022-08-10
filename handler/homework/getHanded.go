package homework

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/file"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary Get handed homework
// @Tags homework
// @Description Get the submitted homework by group
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param group_id query integer true "小组id"
// @Param limit query integer true "limit--偏移量指定开始返回记录之前要跳过的记录数 "
// @Param page  query integer true "page--限制指定要检索的记录数 "
// @Success 200 {object} []file.Homework "{"msg":"查看成功"}"
// @Router /homework/handed [get]
func GetHandedHomework(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("group_id"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
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

	homework, num, err := file.GetHomeworkHanded(ID, limit*page, limit)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, map[string]interface{}{
		"homework": homework,
		"num":      num,
	})
}
