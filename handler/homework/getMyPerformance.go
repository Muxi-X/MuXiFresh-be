package homework

import (
	. "github.com/MuXiFresh-be/handler"
	File "github.com/MuXiFresh-be/service/file"
	"github.com/gin-gonic/gin"
)

// @Summary Get Performance
// @Tags homework
// @Description 查看用户所有作业的完成状况
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []PerformanceResponse{} "{"msg":"查看成功"}"
// @Failure 500 {object} errno.Errno "{"msg":"Error occurred while getting url queries."}"
// @Router /homework/performance [get]
func GetMyPerformance(c *gin.Context) {
	email := c.MustGet("email").(string)
	homework, err := File.GetAllMyHomework(email)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}
	published, err := File.GetAllPublished(email)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}
	length := len(published)
	resp := make([]PerformanceResponse, length)
	for i, m := range published {
		status := 0
		for _, n := range homework {
			if m.ID == n.HomeworkID {
				status = 1
			}
			resp[i] = PerformanceResponse{
				ID:      m.ID,
				Title:   m.Title,
				Content: m.Content,
				URL:     m.FileUrl,
				Status:  status,
			}
		}
	}

	SendResponse(c, nil, resp)
}
