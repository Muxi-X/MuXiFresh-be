package user

import (
	"fmt"
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	service "github.com/MuXiFresh-be/service/user"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine()) //, errno.ErrBind)
		fmt.Println(req)
		return
	}

	err1, err2 := service.Register(req)
	if err2 != nil {
		SendBadRequest(c, err1, nil, err2.Error(), GetLine())
	}

	SendResponse(c, nil, "succeed in registration")

}
