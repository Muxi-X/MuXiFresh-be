package user

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Login ... 学生登录
// @Summary login api
// @Description login the student-forum
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param object body loginRequest true "login_request"
// @Success 200 {object} loginResponse
// @Router /auth/login/student [post]
func Login(c *gin.Context) {
	log.Info("student login function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req loginRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造请求给 login
	// loginReq := &pb.StudentLoginRequest{
	// 	StudentId: req.StudentId,
	// 	Password:  req.Password,
	// }

	// 发送请求
	// loginResp, err := service.UserClient.StudentLogin(context.Background(), loginReq)

	// if err != nil {
	// 	parsedErr := errors.Parse(err.Error())
	// 	detail, errr := e.ParseDetail(parsedErr.Detail)
	//
	// 	finalErrno := errno.InternalServerError
	// 	if errr == nil {
	// 		finalErrno = &errno.Errno{
	// 			Code:    detail.Code,
	// 			Message: detail.Cause,
	// 		}
	// 	}
	// 	SendError(c, finalErrno, nil, err.Error(), GetLine())
	// 	return
	// }

	// 构造返回 response
	// resp := studentLoginResponse{
	// 	Token: loginResp.Token,
	// }
	//
	// SendResponse(c, nil, resp)
}
