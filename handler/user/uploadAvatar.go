package user

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service"
	"github.com/gin-gonic/gin"
)

// UploadAvatar ... 上传头像
// @Summary Avatar api
// @Description
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param file formData file true "文件"
// @Success 200 {object} loginResponse
// @Router /user/login [post]
func UploadAvatar(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
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

	SendResponse(c, nil, map[string]string{
		"Url": url,
	})
}
