package handler

import (
	"as/log"
	"as/util"
	"net/http"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Response 请求响应
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
} // @name Response

func GetLine() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "as/handler/handler.go:30"
	}
	return file + ":" + strconv.Itoa(line)
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	message := ""
	if err != nil {
		message = err.Error()
	}
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

func SendError(c *gin.Context, err error, data interface{}, cause string, source string) {
	log.Error("error",
		zap.String("X-Request-Id", util.GetReqID(c)),
		zap.String("cause", cause),
		zap.String("source", source))

	c.JSON(http.StatusBadRequest, Response{
		Code:    400,
		Message: cause,
		Data:    data,
	})
}
