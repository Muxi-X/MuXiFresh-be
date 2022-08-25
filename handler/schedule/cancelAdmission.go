package schedule

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/pkg/errno"
	service "github.com/MuXiFresh-be/service/schedule"
	"github.com/MuXiFresh-be/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary "取消录取"
// @Description "审阅板块取消录取某位学生"
// @Tags schedule
// @Accept json
// @Produce json
// @Param name path string true "name"
// @param Authorization header string true "token 用户令牌"
// @success 200 {object} handler.Response
// @Router /schedule/cancel_admission/:name [put]
func CancelAdmission(c *gin.Context) {
	log.Info("Cancel one student admission function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	name := c.Param("name")

	// if err := model.DB.Self.Table("schedules").Where("name = ? ", name).Updates(map[string]interface{}{"admission_status": 0}).Error; err != nil {
	// 	SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
	// 	return
	// }
	err := service.CancelAdmit(name)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "取消录取成功")
}
