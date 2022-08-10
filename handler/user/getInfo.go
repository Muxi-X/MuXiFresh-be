package user

import (
	. "github.com/MuXiFresh-be/handler"
	U "github.com/MuXiFresh-be/service/user"
	"github.com/gin-gonic/gin"
)

// @Summary GetInfo
// @Description 得到用户所有的个人信息
// @Tags user
// @Accept  json/application
// @Produce  json/application
// @Param Authorization header string true "token"
// @Success 200 {object}  User
// @Router /user/info [get]
func GetInfo(c *gin.Context) {
	email := c.MustGet("email").(string)

	info, err := U.GetInfo(email)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}
	var user = User{
		Id:        info.ID,
		StudentID: info.StudentId,
		Email:     info.Email,
		Name:      info.Name,
		Avatar:    info.Avatar,
		Role:      info.Role,
	}
	SendResponse(c, nil, user)
}
