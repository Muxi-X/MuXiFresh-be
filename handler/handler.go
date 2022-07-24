package handler

import (
	"runtime"
	"strconv"

	"github.com/MuXiFresh-be/log"
	"github.com/MuXiFresh-be/util"
	"net/http"

	"github.com/MuXiFresh-be/pkg/errno"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Response 请求响应
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
} // @name Response

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	log.Info(message, zap.String("X-Request-Id", util.GetReqID(c)))

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendBadRequest(c *gin.Context, err error, data interface{}, cause string, source string) {
	code, message := errno.DecodeErr(err)
	log.Error(message,
		zap.String("X-Request-Id", util.GetReqID(c)),
		zap.String("cause", cause),
		zap.String("source", source))

	c.JSON(http.StatusBadRequest, Response{
		Code:    code,
		Message: message + ": " + cause,
		Data:    data,
	})
}

func SendError(c *gin.Context, err error, data interface{}, cause string, source string) {
	code, message := errno.DecodeErr(err)
	log.Error(message,
		zap.String("X-Request-Id", util.GetReqID(c)),
		zap.String("cause", cause),
		zap.String("source", source))

	var responseCode = http.StatusInternalServerError

	if code == http.StatusNotFound {
		responseCode = http.StatusNotFound
	} else if code > 20000 {
		responseCode = http.StatusBadRequest
	}

	c.JSON(responseCode, Response{
		Code:    code,
		Message: message + " " + cause,
		Data:    data,
	})
}

func GetLine() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "github.com/MuXiFresh-be/handler/handler.go:67"
	}
	return file + ":" + strconv.Itoa(line)
}
