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

// @Summary "录取"
// @Description "审阅板块录取某位学生"
// @Tags schedule
// @Accept json
// @Produce json
// @Param name path string true "name"
// @param Authorization header string true "token 用户令牌"
// @success 200 {object} handler.Response
// @Router /schedule/admit/:name [put]
func Admit(c *gin.Context) {
	log.Info("Admit one student function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	name := c.Param("name")

	err := service.Admit(name)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "录取成功")
}
