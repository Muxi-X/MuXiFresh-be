package user

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/pkg/errno"
	service "github.com/MuXiFresh-be/service/user"
	"github.com/MuXiFresh-be/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Login ... 登录
// @Summary login api
// @Description
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param object body loginRequest true "login_request"
// @Success 200 {object} loginResponse
// @Router /user/login [post]
func Login(c *gin.Context) {
	log.Info("student login function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	var req loginRequest
	if err := c.ShouldBind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	var db *model.Database
	db.Init()
	defer db.Close()

	token, err := service.Login(req.Email, req.Password)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}

	resp := loginResponse{
		Token: token,
	}

	SendResponse(c, nil, resp)
}
