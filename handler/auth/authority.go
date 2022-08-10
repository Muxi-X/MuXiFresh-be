package auth

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/handler/user"
	"github.com/MuXiFresh-be/pkg/errno"
	User "github.com/MuXiFresh-be/service/user"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary Get Authority
// @Description 查看所有管理员
// @Tags auth
// @Accept  json/application
// @Produce  json/application
// @Param Authorization header string true  "获取email"
// @Param role query integer true "2--管理员   4--超管"
// @Success 200 {object}  user.AdminResponse
// @Router /auth/administrator [get]
func GetAdministrator(c *gin.Context) {
	role, err := strconv.Atoi(c.DefaultQuery("role", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}
	administrators, err := User.GetAuthority(role)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}
	resp := make([]user.AdminResponse, len(administrators))
	for i, m := range administrators {
		resp[i] = user.AdminResponse{
			ID:        m.ID,
			StudentID: m.StudentId,
			Name:      m.Name,
			Email:     m.Email,
			Role:      role,
		}
	}
	SendResponse(c, nil, resp)

}
