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

// @Summary Review homework
// @Tags homework
// @Description review the homework handed
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param id query integer true "不同组别的作业，1-设计组 2-产品组 3-安卓组 4-前端组 5-"
// @Param limit query integer true "limit--偏移量指定开始返回记录之前要跳过的记录数 "
// @Param page  query integer true "page--限制指定要检索的记录数 "
// @Success 200 {object} []file.Homework {} "{"msg":"查看成功"}"
// @Failure 500 {object} errno.Errno "{"msg":"Error occurred while getting url queries."}"
// @Router /homework/review [get]
func ReviewHomework(c *gin.Context) {
	log.Info("Idea getIdeaList function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	id := c.DefaultQuery("id", "0")

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

	homework, num, err := file.Review(id, limit*page, limit)

	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, map[string]interface{}{
		"Homework": homework,
		"Num":      num,
	})

}
