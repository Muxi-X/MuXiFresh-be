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
// @Param id query integer true "作业id"
// @Success 200 {object} []file.Homework {} "{"msg":"查看成功"}"
// @Router /homework/review [get]
func ReviewHomework(c *gin.Context) {
	log.Info("Idea getIdeaList function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))

	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	if err := file.Review(id); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "Success")

}
