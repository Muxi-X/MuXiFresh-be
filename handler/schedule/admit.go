package schedule

import (
	"fmt"

	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	service "github.com/MuXiFresh-be/service/schedule"
	"github.com/gin-gonic/gin"
)

// @Summary "录取"
// @Description "输入姓名和第几次录取的-（1-已报名，2-初试过，3-面试过）"
// @Tags schedule
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Success 200
// @Router /schedule/admit [put]
func Admit(c *gin.Context) {
	email := c.MustGet("email").(string)
	if email == "" {
		return
	}
	var group admitRequest
	if err := c.BindJSON(&group); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}
	fmt.Println(group.AdmissionStatus)
	err := service.Admit(group.Name, group.AdmissionStatus)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "录取成功")

}

// // @Summary "录取"
// // @Description "审阅板块录取某位学生"
// // @Tags schedule
// // @Accept json
// // @Produce json
// // @Param name path string true "name"
// // @param Authorization header string true "token 用户令牌"
// // @success 200 {object} handler.Response
// // @Router /schedule/admit/:name [put]
// func Admit(c *gin.Context) {
// 	log.Info("Admit one student function called.", zap.String("X-Request-Id", util.GetReqID(c)))

// 	name := c.Param("name")

// 	err := service.Admit(name)
// 	if err != nil {
// 		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
// 		return
// 	}

// 	SendResponse(c, nil, "录取成功")
// }
