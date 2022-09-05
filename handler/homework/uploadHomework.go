package homework

import (
	. "github.com/MuXiFresh-be/handler"

	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/file"
	File "github.com/MuXiFresh-be/service/file"
	"github.com/MuXiFresh-be/service/schedule"
	"github.com/MuXiFresh-be/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UploadHomework ... 上传作业
// @Summary Upload homework
// @Description 上传完成的作业
// @Tags homework
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param req body HomeworkRequest true "提交作业内容"
// @Success 200 {object} handler.Response
// @Router /homework [post]
func UploadHomework(c *gin.Context) {
	log.Info("Idea getIdeaList function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	email := c.MustGet("email").(string)
	//IntHomework, _ := strconv.Atoi(c.PostForm("homeworkID"))
	//homeworkID := uint(IntHomework)
	//file, fileHeader, err := c.Request.FormFile("homework")
	//if err != nil {
	//	SendError(c, errno.ErrGetFile, nil, err.Error(), GetLine())
	//	return
	//}
	//fileSize := fileHeader.Size
	//id, _ := c.MustGet("id").(uint32)
	//
	//url, err := service.UploadFile(fileHeader.Filename, id, file, fileSize)
	//if err != nil {
	//	SendError(c, errno.ErrUploadFile, nil, err.Error(), GetLine())
	//	return
	//}
	//
	//if err := Service.HandInHomework(url, email, homeworkID); err != nil {
	//	SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
	//	return
	//}

	var homework HomeworkRequest
	if err := c.ShouldBind(&homework); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}
	hw, err := file.HandInHomework(homework.Title, homework.Content, homework.HomeworkID, homework.FileUrl, email)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	//作业是否全部完成
	homeworkAll, err := File.GetAllMyHomework(email)
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
		for _, n := range homeworkAll {
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
	// fmt.Println(length)
	// for i := 0; i < length; i++ {
	// 	fmt.Println(resp[i].Status)
	// }
	for i := 0; i < length; i++ {
		if resp[i].Status == 0 {
			if err := schedule.UpdateSchedule3(email); err != nil {
				SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
				return
			}
			SendResponse(c, nil, hw)
			return
		}
	}
	// fmt.Println(2333)
	if err := schedule.UpdateSchedule2(email); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, hw)
}
