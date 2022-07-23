package schedule

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary "查看进度"
// @Description "进度查询板块呈现的表格"
// @Tags schedule
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} GetInfoReponse
// @Router api/v1/schedule [get]
func ViewOwnSchedule(c *gin.Context) {
	log.Info("User getOwnSchedule function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	id := c.MustGet("userId")

	var schedule GetInfoReponse
	if err := model.DB.Self.Where("id = ? ", id).Find(&schedule).Error; err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, schedule)
}
