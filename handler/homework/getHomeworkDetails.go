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
	email := c.MustGet("email").(string)
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}
	homework, finished, err := file.GetHomeworkDetails(ID, email)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	status := 0
	for _, m := range finished {
		if m.HomeworkID == homework.ID {
			status = 1
		}
	}
	SendResponse(c, nil, PublishedResponse{
		Title:   homework.Title,
		Content: homework.Content,
		GroupID: homework.GroupID,
		FileUrl: homework.FileUrl,
		Status:  status,
	})
}
