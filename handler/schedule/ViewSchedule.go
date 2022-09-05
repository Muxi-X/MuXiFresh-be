package schedule

import (
	"fmt"

	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/form"
	"github.com/MuXiFresh-be/service/schedule"
	U "github.com/MuXiFresh-be/service/user"
	"github.com/MuXiFresh-be/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ScheduleResponse struct {
	Name            string `json:"name"`
	Major           string `json:"major"`
	Group           string `json:"group"`
	FormStatus      int    `json:"form_status"`
	WorkStatus      int    `json:"work_status"`
	AdmissionStatus int    `json:"admission_status"`
}

// @Summary "查看进度"
// @Description "进度查询板块呈现的表格"
// @Tags schedule
// @Accept json
// @Produce json
// @param Authorization header string true "token 用户令牌"
// @Success 200 {object} schedule.ScheduleModel
// @Router /schedule [get]
func ViewOwnSchedule(c *gin.Context) {
	log.Info("User getOwnSchedule function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	email := c.MustGet("email").(string)

	// var schedule GetInfoResponse
	// if err := model.DB.Self.Table("schedules").Where("email = ? ", email).Find(&schedule).Error; err != nil {
	// 	SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
	// 	return
	// }
	sche, err := schedule.ViewOwn(email)
	Form, err := form.ViewForm1(email)
	fmt.Println("123", Form.Group)
	User, err := U.GetInfo(email)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, ScheduleResponse{
		//Name:            Form.Name,
		Name:            User.Name,
		Major:           User.Major,
		Group:           Form.Group,
		FormStatus:      sche.FormStatus,
		WorkStatus:      sche.WorkStatus,
		AdmissionStatus: sche.AdmissionStatus,
	})
}
