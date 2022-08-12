package homework

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/file"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary Get homework details
// @Tags homework
// @Description  get the published homework details
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param id path integer true "发布的作业id"
// @Success 200 {object} file.HomeworkPublished "{"msg":"查看成功"}"
// @Router /homework/published/details/:id [get]
func GetHomeworkDetails(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}
	homework, err := file.GetHomeworkDetails(ID)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, homework)
}
