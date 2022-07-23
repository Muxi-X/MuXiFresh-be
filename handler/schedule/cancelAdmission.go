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

// @Summary "取消录取"
// @Description "审阅板块取消录取某位学生"
// @Tags schedule
// @Accept json
// @Produce json
// @Param id path int true "target_id"
// @Param token header string true "token"
// @Success 200
// @Router api/v1/schedule [put]
func CancelAdmission(c *gin.Context) {
	log.Info("Cancel one student admission function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	id := c.Param("target_id")

	if err := model.DB.Self.Where("id = ? ", id).Updates(map[string]interface{}{"admission_status": 0}).Error; err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "取消录取成功")
}
