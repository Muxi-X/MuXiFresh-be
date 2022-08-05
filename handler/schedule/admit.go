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

// @Summary "录取"
// @Description "审阅板块录取某位学生"
// @Tags schedule
// @Accept json
// @Produce json
// @Param id path int true "target_id"
// @Param token header string true "token"
// @Success 200
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router api/v1/schedule/admit/:name [put]
func Admit(c *gin.Context) {
	log.Info("Admit one student function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	name := c.Param("name")

	if err := model.DB.Self.Table("schedules").Where("name = ? ", name).Updates(map[string]interface{}{"admission_status": 1}).Error; err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "录取成功")
}
