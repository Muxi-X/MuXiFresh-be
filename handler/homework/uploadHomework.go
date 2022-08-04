package homework

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/file"
	"github.com/MuXiFresh-be/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UploadHomework ... 上传作业
// @Summary Homework api
// @Description
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

	if err := file.HandInHomework(homework.Title, homework.Content, homework.HomeworkID, homework.FileUrl, email); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, "Success")
}
