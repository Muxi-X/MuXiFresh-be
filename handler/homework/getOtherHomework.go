package homework

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/log"
	File "github.com/MuXiFresh-be/service/file"
	"github.com/MuXiFresh-be/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary Get Others' Homework
// @Tags homework
// @Description 查看某人的作业
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param email path integer true "用户的email"
// @Success 200 {object} []file.Homework {} "{"msg":"查看成功"}"
// @Failure 500 {object} errno.Errno "{"msg":"Error occurred while getting url queries."}"
// @Router /homework/:email [get]
func GetOtherHomework(c *gin.Context) {
	log.Info("Idea getIdeaList function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	email := c.Param("email")
	homework, err := File.GetAllMyHomework(email)
	if err != nil {
		SendError(c, err, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, nil, homework)
}
