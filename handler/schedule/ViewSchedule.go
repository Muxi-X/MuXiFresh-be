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
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router api/v1/schedule [get]
func ViewOwnSchedule(c *gin.Context) {
	log.Info("User getOwnSchedule function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	email := c.MustGet("email")

	var schedule GetInfoReponse
	if err := model.DB.Self.Table("schedules").Where("email = ? ", email).Find(&schedule).Error; err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, schedule)
}
