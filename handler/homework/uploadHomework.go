package homework

import (
	. "github.com/MuXiFresh-be/handler"
	File "github.com/MuXiFresh-be/model/file"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UploadHomework ... 上传作业
// @Summary Homework api
// @Description
// @Tags homework
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param homework formData file true "文件"
// @Param id
// @Success 200 {object} loginResponse
// @Router /homework [post]
func UploadHomework(c *gin.Context) {
	email := c.MustGet("email").(string)
	IntHomework, _ := strconv.Atoi(c.PostForm("homeworkID"))
	homeworkID := uint(IntHomework)
	file, fileHeader, err := c.Request.FormFile("homework")
	if err != nil {
		SendError(c, errno.ErrGetFile, nil, err.Error(), GetLine())
		return
	}
	fileSize := fileHeader.Size
	id, _ := c.MustGet("id").(uint32)

	url, err := service.UploadFile(fileHeader.Filename, id, file, fileSize)
	if err != nil {
		SendError(c, errno.ErrUploadFile, nil, err.Error(), GetLine())
		return
	}

	var homework File.Homework
	homework.URL = url
	homework.Email = email
	homework.HomeworkID = homeworkID
	if err := homework.Create(); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, homework)
}
